#!/bin/bash

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Testing gRPC Services ===${NC}\n"

# Test Area Service
echo -e "${GREEN}Testing AreaService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50051 area.AreaService/GetArea
echo -e "\n"

# Test Photo Service
echo -e "${GREEN}Testing PhotoService...${NC}"
echo "Request: {\"id\":100}"
grpcurl -plaintext -d '{"id":100}' localhost:50051 photo.PhotoService/GetPhoto
echo -e "\n"

# Test Restaurant Service
echo -e "${GREEN}Testing RestaurantService...${NC}"
echo "Request: {\"id\":\"0b95fe0d-a323-43d7-954b-2164a1d3242d\"}"
grpcurl -plaintext -d '{"id":"0b95fe0d-a323-43d7-954b-2164a1d3242d"}' localhost:50051 restaurant.RestaurantService/GetRestaurant
echo -e "\n"

# Test Menu Service
echo -e "${GREEN}Testing MenuService...${NC}"
echo "Request: {\"id\":\"116e70bb-c26c-4ec7-8935-7f922e8bf551\"}"
grpcurl -plaintext -d '{"id":"116e70bb-c26c-4ec7-8935-7f922e8bf551"}' localhost:50051 menu.MenuService/GetMenu
echo -e "\n"

# Test Video Service
echo -e "${GREEN}Testing VideoService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50051 video.VideoService/GetVideo
echo -e "\n"

# Test AdminUser Service
echo -e "${GREEN}Testing AdminUserService...${NC}"
echo "Request: {\"id\":\"5a5911ef-e1f9-ea31-a2da-12811e5e843c\"}"
grpcurl -plaintext -d '{"id":"5a5911ef-e1f9-ea31-a2da-12811e5e843c"}' localhost:50051 admin_user.AdminUserService/GetAdminUser
echo -e "\n"

# Test Category Service
echo -e "${GREEN}Testing CategoryService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50051 category.CategoryService/GetCategory
echo -e "\n"

echo -e "${BLUE}=== All tests completed ===${NC}"

