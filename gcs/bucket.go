package gcs
import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
)

type (
	ObjectCallback func(*storage.ObjectHandle) error
	AttrCallback func(*storage.ObjectAttrs) error
)
// cb must be one of ObjectCallback or AttrCallback
//   usage:
//   err := All(ctx,bucket,gcs.ObjectCallback(cb))
//   err := All(ctx,bucket,gcs.AttrCallback(cb))
func All(ctx context.Context, bucket *storage.BucketHandle,	cb interface{}) error {
	var ea func(string,interface{}) error
	ea = func(prefix string,cb1 interface{}) error{
		query := &storage.Query{Prefix:prefix,Delimiter:"/"}
		var pf []string
//		for query != nil {
		objs,err := bucket.List(ctx,query)
		if err != nil{ return err}
		pf=objs.Prefixes
		switch cb2 := cb1.(type){
		case ObjectCallback:
			for _,res := range objs.Results {
				err = cb2(bucket.Object(res.Name))
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

func List(ctx context.Context,
	bucket *storage.BucketHandle,
	prefix string,
	cb func([]string,[]*storage.ObjectAttrs)error)error{
	query := &storage.Query{Prefix:prefix,Delimiter:"/"}
	objs,err := bucket.List(ctx,query)
	if err != nil { return err}
	return cb(objs.Prefixes,objs.Results)
}
