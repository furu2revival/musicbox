# infra

ここではインフラのリソースファイルを管理しています。

## Prerequisites

- AWS
- Terraform
- terraform-docs
- tflint

## ファイルストラクチャ

[Terraform を使用するためのベストプラクティス](https://cloud.google.com/docs/terraform/best-practices-for-terraform) に従った構成になっています。

```
-- modules/
   -- <service-name>/
      -- main.tf
      -- variables.tf
      -- outputs.tf
      -- provider.tf
      -- header.md
      -- README.md
      -- .terraform-docs.yml
   -- ...others...
-- environments/
   -- dev/
      -- main.tf
      -- backend.tf
      -- secrets.yaml
      -- secrets.yaml.encrypted
   -- prod/
      -- main.tf
      -- backend.tf
      -- secrets.yaml
      -- secrets.yaml.encrypted
```

## Secrets の管理方法

```shell
$ cd environments/<environment name>

# 暗号化
$ aws kms encrypt --key-id <KMSのキーID> \
    --plaintext fileb://secrets.yaml \
    --output text \
    --query CiphertextBlob > secrets.yaml.encrypted
# 復号
$ aws kms decrypt \
    --ciphertext-blob fileb://<(cat secrets.yaml.encrypted | base64 -d) \
    --output text \
    --query Plaintext | base64 --decode > secrets.yaml
```

## Terraform ドキュメントの自動生成

### 初回のみ

- `modules/<service-name>/header.md` にタイトルなどヘッダー情報を記載する。
- `modules/<service-name>/.terraform-docs.yml` という名前で以下のファイルを用意する。

```yaml
formatter: markdown table
header-from: header.md
output:
  file: README.md
  mode: inject
  template: |-
    <!-- BEGIN_TF_DOCS -->
    {{ .Content }}
    <!-- END_TF_DOCS -->
```

### ドキュメントの生成

```shell
terraform-docs modules/<service-name>
```
