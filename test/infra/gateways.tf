resource "aws_nat_gateway" "nat_az_a" {
  subnet_id = aws_subnet.public_subnet_az_a.id
  allocation_id = aws_eip.nat_eip_az_a.id
  tags = {
    Name = "nat-az-a"
    ManagedBy = "terraform"
  }
}

resource "aws_nat_gateway" "nat_az_c" {
  subnet_id = aws_subnet.public_subnet_az_c.id
  allocation_id = aws_eip.nat_eip_az_c.id
  tags = {
    Name = "nat-az-c"
    ManagedBy = "terraform"
  }
}
