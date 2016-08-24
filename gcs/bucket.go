package gcs
import (
	"fmt"

	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
)

type (
	ObjectCallback func(*storage.ObjectHandle) error
	AttrCallback func(*storage.ObjectAttrs) error
	Bucket struct{
		B *storage.BucketHandle
		C context.Context
	}
)
// cb must be one of ObjectCallback or AttrCallback
//   usage:
//   err := All(ctx,bucket,gcs.ObjectCallback(cb))
//   err := All(ctx,bucket,gcs.AttrCallback(cb))
func (bucket *Bucket)All(cb interface{}) error {
	var ea func(string,interface{}) error
	ea = func(prefix string,cb1 interface{}) error{
		query := &storage.Query{Prefix:prefix,Delimiter:"/"}
		var pf []string
//		for query != nil {
		objs,err := bucket.B.List(bucket.C,query)
		if err != nil{ return err}
		pf=objs.Prefixes
		switch cb2 := cb1.(type){
		case ObjectCallback:
			for _,res := range objs.Results {
				err = cb2(bucket.B.Object(res.Name))
				if err != nil { return err }
			}
		case AttrCallback:
			for _,res := range objs.Results {
				err = cb2(res)
				if err != nil { return err }
			}
		default:
			return fmt.Errorf("Error: Unsupport callback type %T", cb2 )
		}

		// not sure what is objs.Next, studying it in future
//			query = objs.Next
//		}
		for _,p:=range pf {
			err := ea(p,cb1)
			if err != nil { return err }
		}
		return nil
	}
	return ea("",cb)
}

func (bucket *Bucket)List(prefix string,cb func([]string,[]*storage.ObjectAttrs)error)error{
	query := &storage.Query{Prefix:prefix,Delimiter:"/"}
	objs,err := bucket.B.List(bucket.C,query)
	if err != nil { return err}
	if len(objs.Prefixes)==0 && len(objs.Results)==0 {
		return storage.ErrObjectNotExist
	}
	return cb(objs.Prefixes,objs.Results)
}
// if name of object exists, it will be overwrite
// guessMimeType guess mime type according the filename
func (bucket *Bucket)NewObjectWriter(name string,guessMimeType func(string)string,f func(*storage.Writer)error) error{
	obj := bucket.B.Object(name)
	w := obj.NewWriter(bucket.C)

	if err := f(w); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	if guessMimeType != nil {
		if _,err :=obj.Update(bucket.C,storage.ObjectAttrs{ContentType:guessMimeType(name)}); err != nil {
			return err
		}
	}
	return nil
}
func (bucket *Bucket)NewObjectReader(name string, f func(*storage.Reader)error)error{
	
	r,err:=bucket.B.Object(name).NewReader(bucket.C)
	if err != nil { return err }
	defer r.Close()
	err = f(r)
	if err != nil { return err }

	return nil
}

func (bucket *Bucket)Delete(name string,fb func(string))error{
	err := bucket.B.Object(name).Delete(bucket.C)
	if err != nil {
		// name is perhps prefix
		return bucket.DeletePrefix(name,fb)
	}
	fb(name)
	return nil
}
func (bucket *Bucket)DeletePrefix(prefix string,fb func(string)) error{
	
	return bucket.List(prefix,func(prefixes []string, attrs []*storage.ObjectAttrs)error{

		for _,p:=range prefixes{
			err := bucket.DeletePrefix(p,fb)
			if err != nil {
				return err
			}
		}
		for _,attr := range attrs{
			err := bucket.B.Object(attr.Name).Delete(bucket.C)
			if err != nil {
				return err
			}

			fb(attr.Name)
		}
		return nil
	})
}
