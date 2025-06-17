resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support = true
  tags = {
    Name = "vpc"
    ManagedBy = "terraform"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "igw"
    ManagedBy = "terraform"
  }
}

resource "aws_subnet" "private_subnet_az_a" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "ap-northeast-2a"
  map_public_ip_on_launch = false
  tags = {
    Name = "private-subnet-az-a"
    ManagedBy = "terraform"
  }
}

resource "aws_subnet" "private_subnet_az_c" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "ap-northeast-2c"
  map_public_ip_on_launch = false
  tags = {
    Name = "private-subnet-az-c"
    ManagedBy = "terraform"
  }
}

resource "aws_subnet" "public_subnet_az_a" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.3.0/24"
  availability_zone = "ap-northeast-2a"
  map_public_ip_on_launch = true
  tags = {
    Name = "public-subnet-az-a"
    ManagedBy = "terraform"
  }
}

resource "aws_subnet" "public_subnet_az_c" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.4.0/24"
  availability_zone = "ap-northeast-2c"
  map_public_ip_on_launch = true
  tags = {
    Name = "public-subnet-az-c"
    ManagedBy = "terraform"
  }
}

