package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretConfiguration struct {
	Uri                 *string `json:"uri,omitempty"`
	WorkspaceSecretName *string `json:"workspaceSecretName,omitempty"`
}
