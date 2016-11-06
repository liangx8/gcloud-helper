package gcs
import (
	"fmt"

	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type (
	ObjectCallback func(*storage.ObjectHandle) error
	AttrCallback func(*storage.ObjectAttrs) error
	Bucket struct{
		bh *storage.BucketHandle
		c context.Context
		client *storage.Client
	}
)
// cb must be one of ObjectCallback or AttrCallback
//   usage:
//   err := All(ctx,bucket,gcs.ObjectCallback(cb))
//   err := All(ctx,bucket,gcs.AttrCallback(cb))
func (bucket *Bucket)Objects(cb interface{},query *storage.Query) error {
	itr := bucket.bh.Objects(bucket.c,query)
	for{
		objAttrs,err := itr.Next()
		if err == iterator.Done {
			break;
		}
		if err != nil {
			return err
		}
		switch cb1 :=cb.(type){
		case ObjectCallback:
		case AttrCallback:
			err = cb1(objAttrs)
			if err != nil { return err }
		default:
			return fmt.Errorf("Error: Unspupport callback type %T",cb)
		}
	}
	return nil
}
func (bucket *Bucket)Close() error{
	if bucket.client != nil {
		return bucket.client.Close()
	}
	return nil
}
func (bucket *Bucket)Object(name string) *storage.ObjectHandle{
	return bucket.bh.Object(name)
}
func AllBucket(ctx context.Context,
	client *storage.Client,
	projecdtID string,
	cb func(bkt *storage.BucketAttrs) error) error{
	itr := client.Buckets(ctx,projecdtID)
	for {
		bktAttrs,err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {	return err }
		err = cb(bktAttrs)
		if err != nil { return err }
	}
	return nil
}

// the bucket of blkName must exists,
func NewBucket(ctx context.Context, prjId,blkName string) (*Bucket,error){
	cli,err := storage.NewClient(ctx)
	if err != nil { return nil,err }
	bkt := cli.Bucket(blkName)
	return &Bucket{bh:bkt,c:ctx,client:cli},nil
}
