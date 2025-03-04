# Table: aws_lightsail_container_service_deployments

This table shows data for Lightsail Container Service Deployments.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerServiceDeployment.html

The composite primary key for this table is (**container_service_arn**, **version**).

## Relations

This table depends on [aws_lightsail_container_services](aws_lightsail_container_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|container_service_arn (PK)|`utf8`|
|containers|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|public_endpoint|`json`|
|state|`utf8`|
|version (PK)|`int64`|