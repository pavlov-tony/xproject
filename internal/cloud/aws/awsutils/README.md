Integration tests should be done for this module. To use it you need
some test account.

See redacted sample test in [./utils_test.go]

Expected output :

```
=== RUN   TestFindingReport
manifest path: reportprefix/test-somereport/20180801-20180901/test-somereport-Manifest.json
--- PASS: TestFindingReport (3.33s)
	utils_test.go:24: manifest: {
		  AdditionalArtifacts: [REDSHIFT,QUICKSIGHT],
		  AdditionalSchemaElements: [RESOURCES],
		  Compression: GZIP,
		  Format: textORcsv,
		  ReportName: "test-somereport",
		  S3Bucket: "user-xproject-test-bucket-12345678",
		  S3Prefix: "reportprefix",
		  S3Region: eu-central-1,
		  TimeUnit: HOURLY
		}
	utils_test.go:31: report: reportprefix/test-somereport/20180801-20180901/zxc12345-1234-5678-9012-abcdefghijkl/test-somereport-1.csv.gz
=== RUN   TestGetMostRecentReportInterval
--- PASS: TestGetMostRecentReportInterval (1.66s)
	utils_test.go:50: 20180801-20180901
PASS
ok  	github.com/pavlov-tony/xproject/internal/cloud/aws/awsutils	5.005s

Compilation finished at Sat Sep 22 16:44:33

```
