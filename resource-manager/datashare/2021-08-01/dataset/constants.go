package dataset

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetKind string

const (
	DataSetKindAdlsGenOneFile               DataSetKind = "AdlsGen1File"
	DataSetKindAdlsGenOneFolder             DataSetKind = "AdlsGen1Folder"
	DataSetKindAdlsGenTwoFile               DataSetKind = "AdlsGen2File"
	DataSetKindAdlsGenTwoFileSystem         DataSetKind = "AdlsGen2FileSystem"
	DataSetKindAdlsGenTwoFolder             DataSetKind = "AdlsGen2Folder"
	DataSetKindBlob                         DataSetKind = "Blob"
	DataSetKindBlobFolder                   DataSetKind = "BlobFolder"
	DataSetKindContainer                    DataSetKind = "Container"
	DataSetKindKustoCluster                 DataSetKind = "KustoCluster"
	DataSetKindKustoDatabase                DataSetKind = "KustoDatabase"
	DataSetKindKustoTable                   DataSetKind = "KustoTable"
	DataSetKindSqlDBTable                   DataSetKind = "SqlDBTable"
	DataSetKindSqlDWTable                   DataSetKind = "SqlDWTable"
	DataSetKindSynapseWorkspaceSqlPoolTable DataSetKind = "SynapseWorkspaceSqlPoolTable"
)

func PossibleValuesForDataSetKind() []string {
	return []string{
		string(DataSetKindAdlsGenOneFile),
		string(DataSetKindAdlsGenOneFolder),
		string(DataSetKindAdlsGenTwoFile),
		string(DataSetKindAdlsGenTwoFileSystem),
		string(DataSetKindAdlsGenTwoFolder),
		string(DataSetKindBlob),
		string(DataSetKindBlobFolder),
		string(DataSetKindContainer),
		string(DataSetKindKustoCluster),
		string(DataSetKindKustoDatabase),
		string(DataSetKindKustoTable),
		string(DataSetKindSqlDBTable),
		string(DataSetKindSqlDWTable),
		string(DataSetKindSynapseWorkspaceSqlPoolTable),
	}
}

func parseDataSetKind(input string) (*DataSetKind, error) {
	vals := map[string]DataSetKind{
		"adlsgen1file":                 DataSetKindAdlsGenOneFile,
		"adlsgen1folder":               DataSetKindAdlsGenOneFolder,
		"adlsgen2file":                 DataSetKindAdlsGenTwoFile,
		"adlsgen2filesystem":           DataSetKindAdlsGenTwoFileSystem,
		"adlsgen2folder":               DataSetKindAdlsGenTwoFolder,
		"blob":                         DataSetKindBlob,
		"blobfolder":                   DataSetKindBlobFolder,
		"container":                    DataSetKindContainer,
		"kustocluster":                 DataSetKindKustoCluster,
		"kustodatabase":                DataSetKindKustoDatabase,
		"kustotable":                   DataSetKindKustoTable,
		"sqldbtable":                   DataSetKindSqlDBTable,
		"sqldwtable":                   DataSetKindSqlDWTable,
		"synapseworkspacesqlpooltable": DataSetKindSynapseWorkspaceSqlPoolTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSetKind(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"moving":    ProvisioningStateMoving,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
