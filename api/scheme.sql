CREATE TYPE platform AS ENUM ('aws', 'gcp');
CREATE TYPE service AS ENUM (
	'ec2', 'vpc', 's3', 'rds', 'sns', 'cloudwatch', 'autoscaling', 'route53', 'lambda',
	'app-engine', 'compute-engine', 'kubernetes-engine', 'storage', 'big-query'
);
CREATE TYPE storage AS ENUM ('hdd', 'ssd');
CREATE TYPE term AS ENUM ('hourly', 'daily', 'weekly', 'monthly');
CREATE TYPE region AS ENUM (
	'us-east-1', 'us-east-2', 'us-east-4', 'us-central-1', 'us-west-1', 'us-west-2',
	'sa-east-1',
	'eu-central-1', 'eu-west-1', 'eu-west-2', 'eu-west-3', 'eu-west-4',
	'ap-northeast-1', 'ap-northeast-2', 'ap-northeast-3', 'ap-southeast-1', 'ap-southeast-2', 'ap-south-1'
);

CREATE SCHEMA xproject;

CREATE TABLE xproject.instances (
	id serial PRIMARY KEY,
	provider platform NOT NULL,
	type service NOT NULL,
	core real,
	ram real,
	disk real,
	disk_type storage,
	price_per_month real,
	price_per_hour real,
	lease_type term,
	location region
);
