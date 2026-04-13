package user

import (
	"context"

	"advanced/constants"
	"advanced/proto"
	"advanced/structs"
)

// DeliveryGRPC menangani request dari layanan luar via RPC
type DeliveryGRPC struct {
	proto.UnimplementedUserServiceServer
	usecase IUserUsecase
}

func ProvideDeliveryGRPC(u IUserUsecase) *DeliveryGRPC {
	return &DeliveryGRPC{usecase: u}
}

func (s *DeliveryGRPC) GetUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	usr, err := s.usecase.GetUser(int(req.GetId()), req.GetName(), req.GetEmail())
	if err != nil {
		// Mengembalikan sukses gRPC namun merepresentasikan DB/Application Error
		return &proto.UserResponse{
			Rc:  constants.RCDBError,
			Msg: constants.MsgGagalAmbilUser,
		}, nil
	}
	return &proto.UserResponse{
		Rc:  constants.RCSukses,
		Msg: constants.MsgSuksesAmbilData,
		Data: &proto.UserData{
			Id:       int32(usr.ID),
			Name:     usr.Name,
			Email:    usr.Email,
			Password: "HIDDEN",
		},
	}, nil
}

// RegisterUser (Public API) menerima pendaftaran
func (s *DeliveryGRPC) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.UserResponse, error) {
	user := structs.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.usecase.RegisterUser(&user); err != nil {
		return &proto.UserResponse{
			Rc:  constants.RCDBError,
			Msg: constants.MsgGagalSimpanUser,
		}, nil
	}

	return &proto.UserResponse{
		Rc:  constants.RCSukses,
		Msg: constants.MsgSuksesBikinData,
		Data: &proto.UserData{
			Id:       int32(user.ID),
			Name:     user.Name,
			Email:    user.Email,
			Password: "HIDDEN",
		},
	}, nil
}

// GetAllUsers (Secured API) mengambil seluruh array data user
func (s *DeliveryGRPC) GetAllUsers(ctx context.Context, req *proto.PaginationRequest) (*proto.UserListResponse, error) {
	limit := int(req.GetLimit())
	page := int(req.GetPage())

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	users, err := s.usecase.GetAllUsers(limit, page)
	if err != nil {
		return &proto.UserListResponse{
			Rc:  constants.RCDBError,
			Msg: constants.MsgGagalAmbilUser,
		}, nil
	}

	var protoUsers []*proto.UserData
	for _, u := range users {
		protoUsers = append(protoUsers, &proto.UserData{
			Id:       int32(u.ID),
			Name:     u.Name,
			Email:    u.Email,
			Password: "HIDDEN",
		})
	}

	return &proto.UserListResponse{
		Rc:  constants.RCSukses,
		Msg: constants.MsgSuksesAmbilData,
		Data: &proto.PaginationData{
			Limit: int32(limit),
			Page:  int32(page),
			Data:  protoUsers,
		},
	}, nil
}
