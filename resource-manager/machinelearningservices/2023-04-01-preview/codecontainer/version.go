package codecontainer

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2023-04-01-preview"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/codecontainer/%s", defaultApiVersion)
}
