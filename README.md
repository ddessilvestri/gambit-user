# AWS RDS Setup for E-commerce Project

This guide explains how to set up a **public MySQL RDS instance on AWS**, connect it to a Lambda function triggered by **Cognito PostConfirmation**, and keep costs minimal using the free tier.

---

## 🗂️ Table of Contents

- [1. RDS Database Setup](#1-rds-database-setup)
- [2. VPC and Subnets](#2-vpc-and-subnets)
- [3. DB Subnet Group](#3-db-subnet-group)
- [4. Security Groups](#4-security-groups)
- [5. Secrets Manager](#5-secrets-manager)
- [6. Lambda Function](#6-lambda-function)
- [7. Cognito Integration](#7-cognito-integration)
- [8. Troubleshooting](#8-troubleshooting)

---

## 1. RDS Database Setup

- Go to **AWS RDS > Create Database**
- Choose:
  - Standard Create
  - Engine: MySQL (not Aurora)
  - Template: Free Tier
  - DB instance: `db.t3.micro` or `db.t4g.micro`
- Set:
  - Username: `admin`
  - Initial DB name: `ecommerce`
- Uncheck: Enable Encryption
- **Public access**: ✅ Yes
- Choose existing **DB Subnet Group** with **only public subnets**
- Disable backups and auto minor upgrades (to reduce cost)

---

## 2. VPC and Subnets

- Use your existing VPC
- You need **2 public subnets** with:
  - Route: `0.0.0.0/0` → Internet Gateway
- Optional: Also keep 2 private subnets (for future backend or RDS-private deployment)

---

## 3. DB Subnet Group

- Go to RDS > Subnet Groups > Create
- Name: `public-subnet-group`
- Add only the 2 **public** subnets from your VPC
- Select this subnet group when creating the DB

---

## 4. Security Groups

- Create or edit a security group with:
  - Inbound Rule:
    - Type: MySQL/Aurora
    - Port: 3306
    - Source: `My IP` (for local dev) or `0.0.0.0/0` (testing only, not production)

Attach this security group to your RDS instance.

---

## 5. Secrets Manager

- Go to AWS Secrets Manager > Store a new secret
- Choose: RDS credentials
- Enter:
  - Username: `admin`
  - Password: Your DB password
  - Database name, host, port (3306), engine: `mysql`
- Save the secret (e.g., `rds/ecommerce-admin`)

Your Lambda will access this secret securely.

---

## 6. Lambda Function

- Runtime: `provided.al2` (if using Go custom runtime)
- Handler: `bootstrap`
- Role: Needs access to:
  - `secretsmanager:GetSecretValue`
  - `logs:*` (for CloudWatch)
- VPC: ❌ **No VPC** if your RDS is public (important!)

### Build Go Executable (PowerShell)

```ps1
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o bootstrap main.go
Compress-Archive -Path ./bootstrap -DestinationPath function.zip
