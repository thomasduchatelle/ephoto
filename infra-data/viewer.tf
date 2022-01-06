resource "aws_ssm_parameter" "iam_policy_bucket_ro" {
  name  = "/dphoto/${var.environment_name}/iam/policies/storageROArn"
  type  = "String"
  value = aws_iam_policy.storage_ro.arn
  tags = local.tags
}

resource "aws_ssm_parameter" "iam_policy_dyn_table_rw" {
  name  = "/dphoto/${var.environment_name}/iam/policies/indexRWArn"
  type  = "String"
  value = aws_iam_policy.index_rw.arn
  tags = local.tags
}

resource "aws_ssm_parameter" "storage_bucket_name" {
  name  = "/dphoto/${var.environment_name}/storage/bucketName"
  type  = "String"
  value = aws_s3_bucket.storage.bucket
  tags = local.tags
}

resource "aws_ssm_parameter" "catalog_table_name" {
  name  = "/dphoto/${var.environment_name}/catalog/tableName"
  type  = "String"
  value = local.dynamodb_table_name
  tags = local.tags
}
