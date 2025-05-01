resource "aws_eip" "nat_eip_az_a" {
  tags = {
    Name = "nat-eip-az-a"
    ManagedBy = "terraform"
  }
}

resource "aws_eip" "nat_eip_az_c" {
  tags = {
    Name = "nat-eip-az-c"
    ManagedBy = "terraform"
  }
}

