// Credit for The NATS.IO Authors
// Copyright 2021-2022 The Memphis Authors
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package models
package models

import "time"

type SystemComponent struct {
	Component   string `json:"component"`
	DesiredPods int    `json:"desired_pods"`
	ActualPods  int    `json:"actual_pods"`
}

type MainOverviewData struct {
	TotalStations    int               `json:"total_stations"`
	TotalMessages    int               `json:"total_messages"`
	SystemComponents []SystemComponent `json:"system_components"`
	Stations         []ExtendedStation `json:"stations"`
}

type StationOverviewData struct {
	ConnectedProducers    []ExtendedProducer           `json:"connected_producers"`
	DisconnectedProducers []ExtendedProducer           `json:"disconnected_producers"`
	DeletedProducers      []ExtendedProducer           `json:"deleted_producers"`
	ConnectedCgs          []Cg                         `json:"connected_cgs"`
	DisconnectedCgs       []Cg                         `json:"disconnected_cgs"`
	DeletedCgs            []Cg                         `json:"deleted_cgs"`
	TotalMessages         int                          `json:"total_messages"`
	AvgMsgSize            int64                        `json:"average_message_size"`
	AuditLogs             []AuditLog                   `json:"audit_logs"`
	Messages              []MessageDetails             `json:"messages"`
	PoisonMessages        []LightPoisonMessage         `json:"poison_messages"`
	Tags                  []Tag                        `json:"tags"`
	Leader                string                       `json:"leader"`
	Followers             []string                     `json:"followers"`
	Schema                StationOverviewSchemaDetails `json:"schema"`
}

type GetStationOverviewDataSchema struct {
	StationName string `form:"station_name" json:"station_name"  binding:"required"`
}

type SystemLogsRequest struct {
	LogType  string `form:"log_type" json:"log_type"  binding:"required"`
	StartIdx int    `form:"start_index" json:"start_index"  binding:"required"`
}

type Log struct {
	MessageSeq int       `json:"message_seq"`
	Type       string    `json:"type"`
	Source     string    `json:"source"`
	Data       string    `json:"data"`
	TimeSent   time.Time `json:"creation_date"`
}

type SystemLogsResponse struct {
	Logs []Log `json:"logs"`
}
