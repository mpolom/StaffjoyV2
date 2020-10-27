// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "account.swagger.json",
		FileModTime: time.Unix(1594620771, 0),

		Content: string("{\n  \"swagger\": \"2.0\",\n  \"info\": {\n    \"title\": \"account.proto\",\n    \"version\": \"version not set\"\n  },\n  \"schemes\": [\n    \"http\",\n    \"https\"\n  ],\n  \"consumes\": [\n    \"application/json\"\n  ],\n  \"produces\": [\n    \"application/json\"\n  ],\n  \"paths\": {\n    \"/v1/accounts\": {\n      \"get\": {\n        \"operationId\": \"List\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountAccountList\"\n            }\n          }\n        },\n        \"tags\": [\n          \"AccountService\"\n        ]\n      },\n      \"post\": {\n        \"operationId\": \"Create\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountAccount\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountCreateAccountRequest\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AccountService\"\n        ]\n      }\n    },\n    \"/v1/accounts/{uuid}\": {\n      \"get\": {\n        \"operationId\": \"Get\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountAccount\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"uuid\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\",\n            \"format\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AccountService\"\n        ]\n      },\n      \"put\": {\n        \"operationId\": \"Update\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountAccount\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"uuid\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\",\n            \"format\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountAccount\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AccountService\"\n        ]\n      }\n    },\n    \"/v1/accounts/{uuid}/password\": {\n      \"put\": {\n        \"operationId\": \"UpdatePassword\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/protobufEmpty\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"uuid\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\",\n            \"format\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/accountUpdatePasswordRequest\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AccountService\"\n        ]\n      }\n    }\n  },\n  \"definitions\": {\n    \"accountAccount\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"confirmed_and_active\": {\n          \"type\": \"boolean\",\n          \"format\": \"boolean\"\n        },\n        \"member_since\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"support\": {\n          \"type\": \"boolean\",\n          \"format\": \"boolean\"\n        },\n        \"phonenumber\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"photo_url\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountAccountList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"accounts\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/accountAccount\"\n          }\n        },\n        \"limit\": {\n          \"type\": \"integer\",\n          \"format\": \"int32\"\n        },\n        \"offset\": {\n          \"type\": \"integer\",\n          \"format\": \"int32\"\n        }\n      }\n    },\n    \"accountCreateAccountRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"name\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"phonenumber\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountEmailChangeRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountEmailConfirmation\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountGetAccountByPhonenumberRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"phonenumber\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountGetAccountListRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"offset\": {\n          \"type\": \"integer\",\n          \"format\": \"int32\"\n        },\n        \"limit\": {\n          \"type\": \"integer\",\n          \"format\": \"int32\"\n        }\n      }\n    },\n    \"accountGetAccountRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountGetOrCreateRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"name\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"phonenumber\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountPasswordResetRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountSyncUserRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountTrackEventRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"event\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountUpdatePasswordRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"uuid\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"password\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"accountVerifyPasswordRequest\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        },\n        \"password\": {\n          \"type\": \"string\",\n          \"format\": \"string\"\n        }\n      }\n    },\n    \"protobufEmpty\": {\n      \"type\": \"object\",\n      \"description\": \"service Foo {\\n      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);\\n    }\\n\\nThe JSON representation for `Empty` is empty JSON object `{}`.\",\n      \"title\": \"A generic empty message that you can re-use to avoid defining duplicated\\nempty messages in your APIs. A typical example is to use it as the request\\nor the response type of an API method. For instance:\"\n    }\n  }\n}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1594620771, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "account.swagger.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`swagger`, &embedded.EmbeddedBox{
		Name: `swagger`,
		Time: time.Unix(1594620771, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"account.swagger.json": file2,
		},
	})
}
