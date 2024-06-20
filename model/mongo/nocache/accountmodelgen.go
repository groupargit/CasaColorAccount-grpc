package model

import (
	"context"
	"fmt"

	v1alpha1 "github.com/groupargit/casacoloraccount-grpc/rpc/types/account/v1alpha1"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type accountModel interface {
	LoginUser(ctx context.Context, data *v1alpha1.AccountLoginRequest) (*v1alpha1.AccountLoginResponse, error)
	// InsertCustomer(ctx context.Context, data *v1alpha1.AccountLoginRequest) (*v1alpha1.AccountLoginResponse, error)
	// FindOne(ctx context.Context, id string) (*Customer, error)
	// FindByPod(ctx context.Context, name string) ([]*Customer, error)
	// FindOneByName(ctx context.Context, name string) (*Customer, error)
	// Update(ctx context.Context, data *Customer) error
	// Delete(ctx context.Context, id string) error
}

type defaultAccountModel struct {
	conn *mon.Model
}

func newDefaultAccountModel(conn *mon.Model) *defaultAccountModel {
	return &defaultAccountModel{conn: conn}
}

func (m *defaultAccountModel) LoginUser(ctx context.Context, data *v1alpha1.AccountLoginRequest) (*v1alpha1.AccountLoginResponse, error) {
	//Validate if exist user
	IfExists, err := validateIfExistUserByEmail(ctx, m.conn, data.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("error validate if exist user")
	}
	if !IfExists {
		return nil, fmt.Errorf("error user not exist")
	}
	//Validate password
	//If not exist return error
	//If exist return token
	return nil, nil
}

func validateIfExistUserByEmail(ctx context.Context, conn *mon.Model, email string) (bool, error) {
	var data User
	err := conn.FindOne(ctx, &data, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return true, nil
}

// func (m *defaultPaymentModel) InsertCustomer(ctx context.Context, data *v1alpha1.CreateCustomerRequest) (*v1alpha1.CreateCustomerResponse, error) {
// 	bylogger.LogInfo("CreateCustomerController")
// 	//Validate if exist the organization
// 	//IfExists, _ := validateIfExistOrganizationById(s, req.OrganizationId)
// 	// if err != nil {
// 	// 	bylogs.LogErr("error validate if exist organization")
// 	// 	return nil, fmt.Errorf("error validate if exist organization")
// 	// }
// 	//if !IfExists {
// 	customer_create := Customer{
// 		Id:             primitive.NewObjectID(),
// 		CustomerId:     data.GetCustomerId(),
// 		Name:           data.GetName(),
// 		Email:          data.GetEmail(),
// 		OrganizationId: data.GetOrganizationId(),
// 		Project:        []Project{},
// 		Payment:        []Payment{},
// 	}
// 	//customer_collection := s.Database.Collection("customer_collection")
// 	bylogger.LogInfo("customer_collection", customer_create)
// 	resultInsert, err := m.conn.InsertOne(ctx, customer_create)
// 	if err != nil {
// 		bylogger.LogErr("error insert one", err)
// 		return nil, fmt.Errorf("error insert one")
// 	}

// 	if resultInsert.InsertedID == nil {
// 		bylogger.LogErr("error insert one")
// 		st := status.New(codes.InvalidArgument, "does not save customer")
// 		detail := errdetails.ErrorInfo{
// 			Reason: err.Error(),
// 			Metadata: map[string]string{
// 				"organizationid": data.GetOrganizationId(),
// 				"name":           data.GetName(),
// 				"customer Id":    data.GetCustomerId(),
// 			},
// 		}

// 		st, err = st.WithDetails(&detail)
// 		if err != nil {
// 			bylogger.LogErr(err)
// 			return nil, status.Error(codes.Internal, "does not save customer")
// 		}

// 		return nil, st.Err()
// 	}
// 	response := &v1alpha1.CreateCustomerResponse{
// 		CustomerId: data.CustomerId,
// 	}
// 	return response, nil
// 	// } else {
// 	// 	bylogs.LogErr("error organization already exist")
// 	// 	return nil, status.Error(
// 	// 		codes.Internal,
// 	// 		"error organization already exist",
// 	// 	)
// 	// }

// }

func (m *defaultAccountModel) FindOne(ctx context.Context, id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data User

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAccountModel) Update(ctx context.Context, data *User) error {
	//data.UpdateAt = time.Now()

	_, err := m.conn.ReplaceOne(ctx, bson.M{"_id": data}, data)
	return err
}

func (m *defaultAccountModel) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}

	_, err = m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (m *defaultAccountModel) FindByPod(ctx context.Context, name string) ([]*User, error) {
	var data []*User
	options := options.Find()
	// options.SetSort(bson.M{"created_at": 1})
	options.SetSort(bson.D{{Key: "createAt", Value: 1}})
	options.SetLimit(12)

	err := m.conn.Find(ctx, &data, bson.M{"podname": name}, options)
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAccountModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	var data User

	err := m.conn.FindOne(ctx, &data, bson.M{"name": name})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
