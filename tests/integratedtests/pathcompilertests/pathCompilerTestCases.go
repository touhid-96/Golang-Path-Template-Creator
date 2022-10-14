package pathcompilertests

import (
	"reflect"

	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/enum/osmixtype"
)

var pathCompilerTestCases = []TestWrapper{
	{
		SelectBy:          osmixtype.Unix,
		NameAssert:        "Unix",
		DescriptionAssert: "all unix(ubuntu, debian, linux, darwin ...) related os paths",
		RunsIn: []osmixtype.Variant{
			osmixtype.Windows,
		},
		BaseTestCase: coretests.BaseTestCase{
			Title:        "Unix, Ubuntu, Linux Path should match as per expectation map",
			ArrangeInput: nil,
			ActualInput:  nil,
			ExpectedInput: enumimpl.DynamicMap(map[string]interface{}{
				"AppDbRoot":                "/var/opt/cimux/databases/",
				"ArchiveRoot":              "/var/opt/cimux/archived/",
				"BackupRoot":               "/var/opt/cimux/backups/",
				"CacheTempRoot":            `c:\Windows\Temp\cimux/cache/`,
				"DecompressRoot":           `c:\Windows\Temp\cimux/decompress/`,
				"DefaultConfigFilePath":    "/etc/cimux/config/default-config.json",
				"DefaultEnvPathRoot":       "/var/opt/cimux/env-paths/",
				"DefaultEnvRoot":           "/var/opt/cimux/env/",
				"DefaultInstructionsRoot":  "/var/opt/cimux/instructions/",
				"Description":              "all unix(ubuntu, debian, linux, darwin ...) related os paths",
				"DownloadsRoot":            "/var/opt/cimux/downloads/",
				"EtcAppConfigRoot":         "/etc/cimux/config",
				"EtcAppRoot":               "/etc/cimux",
				"InstructionTempRoot":      `c:\Windows\Temp\cimux/instructions/`,
				"LogAppRoot":               "/var/log/cimux",
				"MigrationCacheRoot":       `c:\Windows\Temp\cimux/migration-cache/`,
				"Name":                     "Unix",
				"PackageTempRoot":          `c:\Windows\Temp\cimux/packages/`,
				"PackagesDownloadRoot":     `/var/opt/cimux/packages-downloaded/`,
				"PackagesRoot":             "/etc/cimux/packages/",
				"PublicRoot":               "/var/www/",
				"ScriptsRoot":              "/var/opt/cimux/scripts/",
				"SnapshotsRoot":            "/var/opt/cimux-snapshots/",
				"SpecificPathFileLocation": "/var/opt/cimux/defined-paths/paths.json",
				"SslRoot":                  "/var/opt/cimux-ssl/",
				"TempRoot":                 `c:\Windows\Temp\cimux`,
				"UserTempRoot":             `c:\Windows\Temp\cimux/users/`,
				"VarAppRoot":               "/var/opt/cimux",
				"VarCacheRoot":             "/var/opt/cimux/cache/",
				"ZipsRoot":                 `/var/opt/cimux/compressed/`,
			}),
			ExpectedTypeOfExpected: reflect.TypeOf(enumimpl.DynamicMap{}),
		},
	},
	{
		RunsIn: []osmixtype.Variant{
			osmixtype.Unix,
		},
		SelectBy:          osmixtype.Unix,
		NameAssert:        "Unix",
		DescriptionAssert: "all unix(ubuntu, debian, linux, darwin ...) related os paths",
		BaseTestCase: coretests.BaseTestCase{
			Title:        "Unix, Ubuntu, Linux Path should match as per expectation map",
			ArrangeInput: nil,
			ActualInput:  nil,
			ExpectedInput: enumimpl.DynamicMap(map[string]interface{}{
				"AppDbRoot":                "/var/opt/cimux/databases/",
				"ArchiveRoot":              "/var/opt/cimux/archived/",
				"BackupRoot":               "/var/opt/cimux/backups/",
				"CacheTempRoot":            "/var/tmp/cimux/cache/",
				"DecompressRoot":           "/var/tmp/cimux/decompress/",
				"DefaultConfigFilePath":    "/etc/cimux/config/default-config.json",
				"DefaultEnvPathRoot":       "/var/opt/cimux/env-paths/",
				"DefaultEnvRoot":           "/var/opt/cimux/env/",
				"DefaultInstructionsRoot":  "/var/opt/cimux/instructions/",
				"Description":              "all unix(ubuntu, debian, linux, darwin ...) related os paths",
				"DownloadsRoot":            "/var/opt/cimux/downloads/",
				"EtcAppConfigRoot":         "/etc/cimux/config",
				"EtcAppRoot":               "/etc/cimux",
				"InstructionTempRoot":      "/var/tmp/cimux/instructions/",
				"LogAppRoot":               "/var/log/cimux",
				"MigrationCacheRoot":       "/var/tmp/cimux/migration-cache/",
				"Name":                     "Unix",
				"PackageTempRoot":          "/var/tmp/cimux/packages/",
				"PackagesDownloadRoot":     `/var/opt/cimux/packages-downloaded/`,
				"PackagesRoot":             "/etc/cimux/packages/",
				"PublicRoot":               "/var/www/",
				"ScriptsRoot":              "/var/opt/cimux/scripts/",
				"SnapshotsRoot":            "/var/opt/cimux-snapshots/",
				"SpecificPathFileLocation": "/var/opt/cimux/defined-paths/paths.json",
				"SslRoot":                  "/var/opt/cimux-ssl/",
				"TempRoot":                 "/var/tmp/cimux",
				"UserTempRoot":             "/var/tmp/cimux/users/",
				"VarAppRoot":               "/var/opt/cimux",
				"VarCacheRoot":             "/var/opt/cimux/cache/",
				"ZipsRoot":                 `/var/opt/cimux/compressed/`,
			}),
			ExpectedTypeOfExpected: reflect.TypeOf(enumimpl.DynamicMap{}),
		},
	},
	{
		RunsIn: []osmixtype.Variant{
			osmixtype.Unix,
		},
		IsTestEnv:         true,
		SelectBy:          osmixtype.Unix,
		NameAssert:        "AnyOs Test",
		DescriptionAssert: "all os test paths",
		BaseTestCase: coretests.BaseTestCase{
			Title:        "Test Env - Unix, Ubuntu, Linux Path should match as per expectation map",
			ArrangeInput: nil,
			ActualInput:  nil,
			ExpectedInput: enumimpl.DynamicMap(map[string]interface{}{
				"AppDbRoot":                "/tmp/cimux-test-env/databases/",
				"ArchiveRoot":              "/tmp/cimux-test-env/archived/",
				"BackupRoot":               "/tmp/cimux-test-env/backups/",
				"CacheTempRoot":            "/tmp/cimux-test-env/tmp//cache/",
				"DecompressRoot":           "/tmp/cimux-test-env/decompress/",
				"DefaultConfigFilePath":    "/tmp/cimux-test-env/config/default-config.json",
				"DefaultEnvPathRoot":       "/tmp/cimux-test-env/env-paths/",
				"DefaultEnvRoot":           "/tmp/cimux-test-env/env/",
				"DefaultInstructionsRoot":  "/tmp/cimux-test-env/instructions/",
				"Description":              "all os test paths",
				"DownloadsRoot":            "/tmp/cimux-test-env/var/opt//downloads/",
				"EtcAppConfigRoot":         "/tmp/cimux-test-env/etc/config",
				"EtcAppRoot":               "/tmp/cimux-test-env/etc/",
				"InstructionTempRoot":      "/tmp/cimux-test-env/tmp//instructions/",
				"LogAppRoot":               "/tmp/cimux-test-env/var/log/",
				"MigrationCacheRoot":       "/tmp/cimux-test-env/tmp//migration-cache/",
				"Name":                     "AnyOs Test",
				"PackageTempRoot":          "/tmp/cimux-test-env/tmp//packages/",
				"PackagesDownloadRoot":     "/tmp/cimux-test-env/packages-downloaded/",
				"PackagesRoot":             "/tmp/cimux-test-env/packages/",
				"PublicRoot":               "/tmp/cimux-test-env/var/www/",
				"ScriptsRoot":              "/tmp/cimux-test-env/scripts/",
				"SnapshotsRoot":            "/tmp/cimux-test-env/all-snapshots/",
				"SpecificPathFileLocation": "/tmp/cimux-test-env/defined-paths/paths.json",
				"SslRoot":                  "/tmp/cimux-test-env/all-ssl/",
				"TempRoot":                 "/tmp/cimux-test-env/tmp/",
				"UserTempRoot":             "/tmp/cimux-test-env/tmp//users/",
				"VarAppRoot":               "/tmp/cimux-test-env",
				"VarCacheRoot":             "/tmp/cimux-test-env/var/opt//cache/",
				"ZipsRoot":                 "/tmp/cimux-test-env/compressed/",
			}),
			ExpectedTypeOfExpected: reflect.TypeOf(enumimpl.DynamicMap{}),
		},
	},
	{
		RunsIn: []osmixtype.Variant{
			osmixtype.Windows,
		},
		IsTestEnv:         true,
		SelectBy:          osmixtype.Windows,
		NameAssert:        "AnyOs Test",
		DescriptionAssert: "all os test paths",
		BaseTestCase: coretests.BaseTestCase{
			Title:        "Test Env - Windows Path should match as per expectation map",
			ArrangeInput: nil,
			ActualInput:  nil,
			ExpectedInput: enumimpl.DynamicMap(map[string]interface{}{
				`AppDbRoot`:                `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/databases/`,
				`ArchiveRoot`:              `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/archived/`,
				`BackupRoot`:               `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/backups/`,
				`CacheTempRoot`:            `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp//cache/`,
				`DecompressRoot`:           `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/decompress/`,
				`DefaultConfigFilePath`:    `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/config/default-config.json`,
				`DefaultEnvPathRoot`:       `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/env-paths/`,
				`DefaultEnvRoot`:           `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/env/`,
				`DefaultInstructionsRoot`:  `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/instructions/`,
				`Description`:              `all os test paths`,
				`DownloadsRoot`:            `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/var/opt//downloads/`,
				`EtcAppConfigRoot`:         `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/etc/config`,
				`EtcAppRoot`:               `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/etc/`,
				`InstructionTempRoot`:      `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp//instructions/`,
				`LogAppRoot`:               `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/var/log/`,
				`MigrationCacheRoot`:       `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp//migration-cache/`,
				`Name`:                     `AnyOs Test`,
				`PackageTempRoot`:          `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp//packages/`,
				`PackagesDownloadRoot`:     `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/packages-downloaded/`,
				`PackagesRoot`:             `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/packages/`,
				"PublicRoot":               `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/var/www/`,
				`ScriptsRoot`:              `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/scripts/`,
				`SnapshotsRoot`:            `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/all-snapshots/`,
				`SpecificPathFileLocation`: `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/defined-paths/paths.json`,
				"SslRoot":                  `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/all-ssl/`,
				`TempRoot`:                 `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp/`,
				`UserTempRoot`:             `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/tmp//users/`,
				`VarAppRoot`:               `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env`,
				`VarCacheRoot`:             `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/var/opt//cache/`,
				`ZipsRoot`:                 `C:\Users\ADMINI~1\AppData\Local\Temp\cimux-test-env/compressed/`,
			}),
			ExpectedTypeOfExpected: reflect.TypeOf(enumimpl.DynamicMap{}),
		},
	},
}
