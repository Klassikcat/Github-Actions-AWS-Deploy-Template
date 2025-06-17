
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }
}

resource "aws_route_table" "private_route_table_az_a" {
  vpc_id = aws_vpc.vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_az_a.id
  }
}

resource "aws_route_table" "private_route_table_az_c" {
  vpc_id = aws_vpc.vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_az_c.id
  }
}

resource "aws_route_table_association" "public_subnet_az_a_association" {
  subnet_id = aws_subnet.public_subnet_az_a.id
  route_table_id = aws_route_table.public_route_table.id
}

resource "aws_route_table_association" "public_subnet_az_c_association" {
  subnet_id = aws_subnet.public_subnet_az_c.id
  route_table_id = aws_route_table.public_route_table.id
}



