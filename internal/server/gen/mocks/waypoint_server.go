// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gen "github.com/hashicorp/waypoint/internal/server/gen"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"
)

// WaypointServer is an autogenerated mock type for the WaypointServer type
type WaypointServer struct {
	mock.Mock
}

// BootstrapToken provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) BootstrapToken(_a0 context.Context, _a1 *emptypb.Empty) (*gen.NewTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.NewTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.NewTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.NewTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelJob provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) CancelJob(_a0 context.Context, _a1 *gen.CancelJobRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gen.CancelJobRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.CancelJobRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertInviteToken provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ConvertInviteToken(_a0 context.Context, _a1 *gen.ConvertInviteTokenRequest) (*gen.NewTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.NewTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ConvertInviteTokenRequest) *gen.NewTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.NewTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ConvertInviteTokenRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHostname provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) CreateHostname(_a0 context.Context, _a1 *gen.CreateHostnameRequest) (*gen.CreateHostnameResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.CreateHostnameResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.CreateHostnameRequest) *gen.CreateHostnameResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.CreateHostnameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.CreateHostnameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSnapshot provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) CreateSnapshot(_a0 *emptypb.Empty, _a1 gen.Waypoint_CreateSnapshotServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*emptypb.Empty, gen.Waypoint_CreateSnapshotServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHostname provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) DeleteHostname(_a0 context.Context, _a1 *gen.DeleteHostnameRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeleteHostnameRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.DeleteHostnameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EntrypointConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) EntrypointConfig(_a0 *gen.EntrypointConfigRequest, _a1 gen.Waypoint_EntrypointConfigServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gen.EntrypointConfigRequest, gen.Waypoint_EntrypointConfigServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntrypointExecStream provides a mock function with given fields: _a0
func (_m *WaypointServer) EntrypointExecStream(_a0 gen.Waypoint_EntrypointExecStreamServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_EntrypointExecStreamServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntrypointLogStream provides a mock function with given fields: _a0
func (_m *WaypointServer) EntrypointLogStream(_a0 gen.Waypoint_EntrypointLogStreamServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_EntrypointLogStreamServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindExecInstance provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) FindExecInstance(_a0 context.Context, _a1 *gen.FindExecInstanceRequest) (*gen.FindExecInstanceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.FindExecInstanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.FindExecInstanceRequest) *gen.FindExecInstanceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.FindExecInstanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.FindExecInstanceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateInviteToken provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GenerateInviteToken(_a0 context.Context, _a1 *gen.InviteTokenRequest) (*gen.NewTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.NewTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.InviteTokenRequest) *gen.NewTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.NewTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.InviteTokenRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateLoginToken provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GenerateLoginToken(_a0 context.Context, _a1 *emptypb.Empty) (*gen.NewTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.NewTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.NewTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.NewTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuild provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetBuild(_a0 context.Context, _a1 *gen.GetBuildRequest) (*gen.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Build
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetBuildRequest) *gen.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetBuildRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetConfig(_a0 context.Context, _a1 *gen.ConfigGetRequest) (*gen.ConfigGetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ConfigGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ConfigGetRequest) *gen.ConfigGetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ConfigGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ConfigGetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigSource provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetConfigSource(_a0 context.Context, _a1 *gen.GetConfigSourceRequest) (*gen.GetConfigSourceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetConfigSourceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetConfigSourceRequest) *gen.GetConfigSourceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetConfigSourceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetConfigSourceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployment provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetDeployment(_a0 context.Context, _a1 *gen.GetDeploymentRequest) (*gen.Deployment, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetDeploymentRequest) *gen.Deployment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetDeploymentRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetJob(_a0 context.Context, _a1 *gen.GetJobRequest) (*gen.Job, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Job
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetJobRequest) *gen.Job); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetJobRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobStream provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetJobStream(_a0 *gen.GetJobStreamRequest, _a1 gen.Waypoint_GetJobStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gen.GetJobStreamRequest, gen.Waypoint_GetJobStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLatestBuild provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetLatestBuild(_a0 context.Context, _a1 *gen.GetLatestBuildRequest) (*gen.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Build
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetLatestBuildRequest) *gen.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetLatestBuildRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestPushedArtifact provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetLatestPushedArtifact(_a0 context.Context, _a1 *gen.GetLatestPushedArtifactRequest) (*gen.PushedArtifact, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.PushedArtifact
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetLatestPushedArtifactRequest) *gen.PushedArtifact); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.PushedArtifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetLatestPushedArtifactRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestRelease provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetLatestRelease(_a0 context.Context, _a1 *gen.GetLatestReleaseRequest) (*gen.Release, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Release
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetLatestReleaseRequest) *gen.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetLatestReleaseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogStream provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetLogStream(_a0 *gen.GetLogStreamRequest, _a1 gen.Waypoint_GetLogStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gen.GetLogStreamRequest, gen.Waypoint_GetLogStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProject provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetProject(_a0 context.Context, _a1 *gen.GetProjectRequest) (*gen.GetProjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetProjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetProjectRequest) *gen.GetProjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetProjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetProjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPushedArtifact provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetPushedArtifact(_a0 context.Context, _a1 *gen.GetPushedArtifactRequest) (*gen.PushedArtifact, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.PushedArtifact
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetPushedArtifactRequest) *gen.PushedArtifact); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.PushedArtifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetPushedArtifactRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRelease provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetRelease(_a0 context.Context, _a1 *gen.GetReleaseRequest) (*gen.Release, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Release
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetReleaseRequest) *gen.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetReleaseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunner provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetRunner(_a0 context.Context, _a1 *gen.GetRunnerRequest) (*gen.Runner, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.Runner
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetRunnerRequest) *gen.Runner); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.Runner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetRunnerRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetServerConfig(_a0 context.Context, _a1 *emptypb.Empty) (*gen.GetServerConfigResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetServerConfigResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.GetServerConfigResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetServerConfigResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatusReport provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetStatusReport(_a0 context.Context, _a1 *gen.GetStatusReportRequest) (*gen.StatusReport, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.StatusReport
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetStatusReportRequest) *gen.StatusReport); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.StatusReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetStatusReportRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersionInfo provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetVersionInfo(_a0 context.Context, _a1 *emptypb.Empty) (*gen.GetVersionInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetVersionInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.GetVersionInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetVersionInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkspace provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) GetWorkspace(_a0 context.Context, _a1 *gen.GetWorkspaceRequest) (*gen.GetWorkspaceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetWorkspaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetWorkspaceRequest) *gen.GetWorkspaceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetWorkspaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetWorkspaceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBuilds provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListBuilds(_a0 context.Context, _a1 *gen.ListBuildsRequest) (*gen.ListBuildsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListBuildsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListBuildsRequest) *gen.ListBuildsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListBuildsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListBuildsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeployments provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListDeployments(_a0 context.Context, _a1 *gen.ListDeploymentsRequest) (*gen.ListDeploymentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListDeploymentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListDeploymentsRequest) *gen.ListDeploymentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListDeploymentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListDeploymentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHostnames provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListHostnames(_a0 context.Context, _a1 *gen.ListHostnamesRequest) (*gen.ListHostnamesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListHostnamesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListHostnamesRequest) *gen.ListHostnamesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListHostnamesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListHostnamesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInstances provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListInstances(_a0 context.Context, _a1 *gen.ListInstancesRequest) (*gen.ListInstancesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListInstancesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListInstancesRequest) *gen.ListInstancesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListInstancesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListInstancesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListProjects(_a0 context.Context, _a1 *emptypb.Empty) (*gen.ListProjectsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListProjectsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.ListProjectsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListProjectsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPushedArtifacts provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListPushedArtifacts(_a0 context.Context, _a1 *gen.ListPushedArtifactsRequest) (*gen.ListPushedArtifactsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListPushedArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListPushedArtifactsRequest) *gen.ListPushedArtifactsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListPushedArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListPushedArtifactsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReleases provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListReleases(_a0 context.Context, _a1 *gen.ListReleasesRequest) (*gen.ListReleasesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListReleasesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListReleasesRequest) *gen.ListReleasesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListReleasesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListReleasesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStatusReports provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListStatusReports(_a0 context.Context, _a1 *gen.ListStatusReportsRequest) (*gen.ListStatusReportsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListStatusReportsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListStatusReportsRequest) *gen.ListStatusReportsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListStatusReportsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListStatusReportsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkspaces provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ListWorkspaces(_a0 context.Context, _a1 *emptypb.Empty) (*gen.ListWorkspacesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListWorkspacesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *gen.ListWorkspacesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListWorkspacesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueJob provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) QueueJob(_a0 context.Context, _a1 *gen.QueueJobRequest) (*gen.QueueJobResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.QueueJobResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.QueueJobRequest) *gen.QueueJobResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.QueueJobResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.QueueJobRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreSnapshot provides a mock function with given fields: _a0
func (_m *WaypointServer) RestoreSnapshot(_a0 gen.Waypoint_RestoreSnapshotServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_RestoreSnapshotServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunnerConfig provides a mock function with given fields: _a0
func (_m *WaypointServer) RunnerConfig(_a0 gen.Waypoint_RunnerConfigServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_RunnerConfigServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunnerGetDeploymentConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) RunnerGetDeploymentConfig(_a0 context.Context, _a1 *gen.RunnerGetDeploymentConfigRequest) (*gen.RunnerGetDeploymentConfigResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.RunnerGetDeploymentConfigResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.RunnerGetDeploymentConfigRequest) *gen.RunnerGetDeploymentConfigResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.RunnerGetDeploymentConfigResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.RunnerGetDeploymentConfigRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunnerJobStream provides a mock function with given fields: _a0
func (_m *WaypointServer) RunnerJobStream(_a0 gen.Waypoint_RunnerJobStreamServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_RunnerJobStreamServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) SetConfig(_a0 context.Context, _a1 *gen.ConfigSetRequest) (*gen.ConfigSetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ConfigSetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ConfigSetRequest) *gen.ConfigSetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ConfigSetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ConfigSetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConfigSource provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) SetConfigSource(_a0 context.Context, _a1 *gen.SetConfigSourceRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SetConfigSourceRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.SetConfigSourceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetServerConfig provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) SetServerConfig(_a0 context.Context, _a1 *gen.SetServerConfigRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SetServerConfigRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.SetServerConfigRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartExecStream provides a mock function with given fields: _a0
func (_m *WaypointServer) StartExecStream(_a0 gen.Waypoint_StartExecStreamServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gen.Waypoint_StartExecStreamServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertApplication provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertApplication(_a0 context.Context, _a1 *gen.UpsertApplicationRequest) (*gen.UpsertApplicationResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertApplicationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertApplicationRequest) *gen.UpsertApplicationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertApplicationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertApplicationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertBuild provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertBuild(_a0 context.Context, _a1 *gen.UpsertBuildRequest) (*gen.UpsertBuildResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertBuildResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertBuildRequest) *gen.UpsertBuildResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertBuildResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertBuildRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDeployment provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertDeployment(_a0 context.Context, _a1 *gen.UpsertDeploymentRequest) (*gen.UpsertDeploymentResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertDeploymentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertDeploymentRequest) *gen.UpsertDeploymentResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertDeploymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertDeploymentRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertProject provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertProject(_a0 context.Context, _a1 *gen.UpsertProjectRequest) (*gen.UpsertProjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertProjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertProjectRequest) *gen.UpsertProjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertProjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertProjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertPushedArtifact provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertPushedArtifact(_a0 context.Context, _a1 *gen.UpsertPushedArtifactRequest) (*gen.UpsertPushedArtifactResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertPushedArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertPushedArtifactRequest) *gen.UpsertPushedArtifactResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertPushedArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertPushedArtifactRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertRelease provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertRelease(_a0 context.Context, _a1 *gen.UpsertReleaseRequest) (*gen.UpsertReleaseResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertReleaseRequest) *gen.UpsertReleaseResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertReleaseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertStatusReport provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) UpsertStatusReport(_a0 context.Context, _a1 *gen.UpsertStatusReportRequest) (*gen.UpsertStatusReportResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpsertStatusReportResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpsertStatusReportRequest) *gen.UpsertStatusReportResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpsertStatusReportResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpsertStatusReportRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateJob provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) ValidateJob(_a0 context.Context, _a1 *gen.ValidateJobRequest) (*gen.ValidateJobResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ValidateJobResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ValidateJobRequest) *gen.ValidateJobResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ValidateJobResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ValidateJobRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaypointHclFmt provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) WaypointHclFmt(_a0 context.Context, _a1 *gen.WaypointHclFmtRequest) (*gen.WaypointHclFmtResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.WaypointHclFmtResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.WaypointHclFmtRequest) *gen.WaypointHclFmtResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.WaypointHclFmtResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.WaypointHclFmtRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// XListJobs provides a mock function with given fields: _a0, _a1
func (_m *WaypointServer) XListJobs(_a0 context.Context, _a1 *gen.ListJobsRequest) (*gen.ListJobsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListJobsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListJobsRequest) *gen.ListJobsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListJobsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListJobsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
