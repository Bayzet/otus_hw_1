// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/calendar.proto

package calendarpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error)
	ListEventsForDay(ctx context.Context, in *ListEventsForDayRequest, opts ...grpc.CallOption) (*Events, error)
	ListEventsForWeek(ctx context.Context, in *ListEventsForWeekRequest, opts ...grpc.CallOption) (*Events, error)
	ListEventsForMonth(ctx context.Context, in *ListEventsForMonthRequest, opts ...grpc.CallOption) (*Events, error)
	FindEventByID(ctx context.Context, in *FindEventByIDRequest, opts ...grpc.CallOption) (*Event, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error) {
	out := new(SuccessOrErrorResponse)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error) {
	out := new(SuccessOrErrorResponse)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*SuccessOrErrorResponse, error) {
	out := new(SuccessOrErrorResponse)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListEventsForDay(ctx context.Context, in *ListEventsForDayRequest, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/ListEventsForDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListEventsForWeek(ctx context.Context, in *ListEventsForWeekRequest, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/ListEventsForWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListEventsForMonth(ctx context.Context, in *ListEventsForMonthRequest, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/ListEventsForMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) FindEventByID(ctx context.Context, in *FindEventByIDRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.v1.Calendar/FindEventByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*SuccessOrErrorResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*SuccessOrErrorResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*SuccessOrErrorResponse, error)
	ListEventsForDay(context.Context, *ListEventsForDayRequest) (*Events, error)
	ListEventsForWeek(context.Context, *ListEventsForWeekRequest) (*Events, error)
	ListEventsForMonth(context.Context, *ListEventsForMonthRequest) (*Events, error)
	FindEventByID(context.Context, *FindEventByIDRequest) (*Event, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) CreateEvent(context.Context, *CreateEventRequest) (*SuccessOrErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarServer) UpdateEvent(context.Context, *UpdateEventRequest) (*SuccessOrErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarServer) DeleteEvent(context.Context, *DeleteEventRequest) (*SuccessOrErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarServer) ListEventsForDay(context.Context, *ListEventsForDayRequest) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsForDay not implemented")
}
func (UnimplementedCalendarServer) ListEventsForWeek(context.Context, *ListEventsForWeekRequest) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsForWeek not implemented")
}
func (UnimplementedCalendarServer) ListEventsForMonth(context.Context, *ListEventsForMonthRequest) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsForMonth not implemented")
}
func (UnimplementedCalendarServer) FindEventByID(context.Context, *FindEventByIDRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEventByID not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListEventsForDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsForDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListEventsForDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/ListEventsForDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListEventsForDay(ctx, req.(*ListEventsForDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListEventsForWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsForWeekRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListEventsForWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/ListEventsForWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListEventsForWeek(ctx, req.(*ListEventsForWeekRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListEventsForMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsForMonthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListEventsForMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/ListEventsForMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListEventsForMonth(ctx, req.(*ListEventsForMonthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_FindEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEventByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).FindEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.v1.Calendar/FindEventByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).FindEventByID(ctx, req.(*FindEventByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.v1.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
		{
			MethodName: "ListEventsForDay",
			Handler:    _Calendar_ListEventsForDay_Handler,
		},
		{
			MethodName: "ListEventsForWeek",
			Handler:    _Calendar_ListEventsForWeek_Handler,
		},
		{
			MethodName: "ListEventsForMonth",
			Handler:    _Calendar_ListEventsForMonth_Handler,
		},
		{
			MethodName: "FindEventByID",
			Handler:    _Calendar_FindEventByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calendar.proto",
}