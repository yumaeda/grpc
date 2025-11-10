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
grpcurl -plaintext -d '{"id":1}' localhost:50050 area.AreaService/GetArea
echo -e "\n"

# Test Photo Service
echo -e "${GREEN}Testing PhotoService...${NC}"
echo "Request: {\"id\":100}"
grpcurl -plaintext -d '{"id":100}' localhost:50050 photo.PhotoService/GetPhoto
echo -e "\n"

# Test Menu Service
echo -e "${GREEN}Testing MenuService...${NC}"
echo "Request: {\"id\":\"116e70bb-c26c-4ec7-8935-7f922e8bf551\"}"
grpcurl -plaintext -d '{"id":"116e70bb-c26c-4ec7-8935-7f922e8bf551"}' localhost:50050 menu.MenuService/GetMenu
echo -e "\n"

# Test AdminUser Service
echo -e "${GREEN}Testing AdminUserService...${NC}"
echo "Request: {\"id\":\"5a5911ef-e1f9-ea31-a2da-12811e5e843c\"}"
grpcurl -plaintext -d '{"id":"5a5911ef-e1f9-ea31-a2da-12811e5e843c"}' localhost:50050 admin_user.AdminUserService/GetAdminUser
echo -e "\n"

# Test Category Service
echo -e "${GREEN}Testing CategoryService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50050 category.CategoryService/GetCategory
echo -e "\n"

# Test Ranking Service - GetRanking
echo -e "${GREEN}Testing RankingService - GetRanking...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50050 ranking.RankingService/GetRanking
echo -e "\n"

# Test Genre Service
echo -e "${GREEN}Testing GenreService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50050 genre.GenreService/GetGenre
echo -e "\n"

# Test Dish Service
echo -e "${GREEN}Testing DishService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50050 dish.DishService/GetDish
echo -e "\n"

# Test Drink Service
echo -e "${GREEN}Testing DrinkService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50050 drink.DrinkService/GetDrink
echo -e "\n"

echo -e "${BLUE}=== All tests completed ===${NC}"
