package mongorepo

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/product/dto"
	"edot/ecommerce/product/entity"
	"edot/ecommerce/product/repository/mongorepo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type product struct {
	c    *mongo.Client
	coll *mongo.Collection
}

func NewProductRepository(c *mongo.Client, coll *mongo.Collection) entity.IProductRepository {
	return &product{
		c, coll,
	}
}

func (r *product) Find(ctx context.Context, input dto.GetListProductInput) (products []entity.Product, err error) {
	filter := bson.D{}
	if input.Keyword != "" {
		filter = bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: input.Keyword}}}}
	}

	sort := bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}
	projection := bson.D{{Key: "name", Value: 1}, {Key: "price", Value: 1}, {Key: "store_id", Value: 1}, {Key: "description", Value: 1}, {Key: "_id", Value: 1}, {Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}
	opts := options.Find().
		SetSort(sort).
		SetProjection(projection).
		SetLimit(int64(input.Limit)).SetSkip(int64((input.Page - 1) * int16(input.Limit)))

	cursor, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}
	m := []model.MProduct{}
	if err = cursor.All(ctx, &m); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	for _, v := range m {
		products = append(products, v.ToEntity())
	}

	return
}

func (r *product) FindProductByIds(ctx context.Context, productIDs []string) (products []entity.Product, err error) {
	ids := make([]primitive.ObjectID, len(productIDs))
	for i, v := range productIDs {
		obj, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "invalid product object id")
			return products, err
		}

		ids[i] = obj
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
	projection := bson.D{{Key: "name", Value: 1}, {Key: "price", Value: 1}, {Key: "store_id", Value: 1}, {Key: "description", Value: 1}, {Key: "_id", Value: 1}}
	opts := options.Find().
		SetProjection(projection)

	cursor, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}
	m := []model.MProduct{}
	if err = cursor.All(ctx, &m); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	for _, v := range m {
		products = append(products, v.ToEntity())
	}

	return
}
