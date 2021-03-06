package redis

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DayOfWeek enumerates the values for day of week.
type DayOfWeek string

const (
	// Everyday ...
	Everyday DayOfWeek = "Everyday"
	// Friday ...
	Friday DayOfWeek = "Friday"
	// Monday ...
	Monday DayOfWeek = "Monday"
	// Saturday ...
	Saturday DayOfWeek = "Saturday"
	// Sunday ...
	Sunday DayOfWeek = "Sunday"
	// Thursday ...
	Thursday DayOfWeek = "Thursday"
	// Tuesday ...
	Tuesday DayOfWeek = "Tuesday"
	// Wednesday ...
	Wednesday DayOfWeek = "Wednesday"
	// Weekend ...
	Weekend DayOfWeek = "Weekend"
)

// PossibleDayOfWeekValues returns an array of possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{Everyday, Friday, Monday, Saturday, Sunday, Thursday, Tuesday, Wednesday, Weekend}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// RebootType enumerates the values for reboot type.
type RebootType string

const (
	// AllNodes ...
	AllNodes RebootType = "AllNodes"
	// PrimaryNode ...
	PrimaryNode RebootType = "PrimaryNode"
	// SecondaryNode ...
	SecondaryNode RebootType = "SecondaryNode"
)

// PossibleRebootTypeValues returns an array of possible values for the RebootType const type.
func PossibleRebootTypeValues() []RebootType {
	return []RebootType{AllNodes, PrimaryNode, SecondaryNode}
}

// ReplicationRole enumerates the values for replication role.
type ReplicationRole string

const (
	// ReplicationRolePrimary ...
	ReplicationRolePrimary ReplicationRole = "Primary"
	// ReplicationRoleSecondary ...
	ReplicationRoleSecondary ReplicationRole = "Secondary"
)

// PossibleReplicationRoleValues returns an array of possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{ReplicationRolePrimary, ReplicationRoleSecondary}
}

// SkuFamily enumerates the values for sku family.
type SkuFamily string

const (
	// C ...
	C SkuFamily = "C"
	// P ...
	P SkuFamily = "P"
)

// PossibleSkuFamilyValues returns an array of possible values for the SkuFamily const type.
func PossibleSkuFamilyValues() []SkuFamily {
	return []SkuFamily{C, P}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Premium, Standard}
}
