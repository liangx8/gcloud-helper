package gcs
import (
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)
func MakeBucket(ctx context.Context,name string) (bkt *storage.BucketHandle,closer func(),err error){
	client,err := storage.NewClient(ctx)
	if err != nil {
		return nil,nil,err
	}
	return client.Bucket(name),func(){
		err := client.Close()
		if err != nil {
			panic(err)
		}
	},nil
}
func Objects(ctx context.Context,bkt *storage.BucketHandle,cb func(*storage.ObjectAttrs) error,prefix string)error {
	query := & storage.Query{
		Delimiter:"/",
		Prefix:prefix,
	}
	itr := bkt.Objects(ctx,query)
	for{
		objAttrs,err := itr.Next()
		if err == iterator.Done {
			break;
		}
		if err != nil {
			return err
		}
		err = cb(objAttrs)
		if err != nil { return err }
	}
	return nil
}

func StringTranslate(f func(string)error)func (*storage.ObjectAttrs)error{
	return func(oa *storage.ObjectAttrs)error{
		return f(oa.Name)
	}
}
