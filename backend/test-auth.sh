#!/bin/bash

# Test Authentication Endpoints

BASE_URL="http://localhost:8080/api/v1"

echo "🧪 Testing Authentication Endpoints"
echo "===================================="
echo ""

# Test 1: Health Check
echo "1. Health Check"
echo "   GET /health"
response=$(curl -s http://localhost:8080/health)
echo "   Response: $response"
echo ""

# Test 2: Register Admin User
echo "2. Register Admin User"
echo "   POST $BASE_URL/auth/register"
register_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "admin123",
    "displayName": "Admin User",
    "role": "admin"
  }' \
  $BASE_URL/auth/register)

echo "   Response: $register_response"
admin_token=$(echo $register_response | jq -r '.token // empty')
echo "   Token: $admin_token"
echo ""

# Test 3: Register Learner User
echo "3. Register Learner User"
echo "   POST $BASE_URL/auth/register"
learner_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "learner@learng.com",
    "password": "learner123",
    "displayName": "Test Learner",
    "role": "learner"
  }' \
  $BASE_URL/auth/register)

echo "   Response: $learner_response"
learner_token=$(echo $learner_response | jq -r '.token // empty')
echo ""

# Test 4: Try to register duplicate user (should fail)
echo "4. Register Duplicate User (should fail)"
echo "   POST $BASE_URL/auth/register"
duplicate_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "password123",
    "displayName": "Duplicate Admin",
    "role": "admin"
  }' \
  $BASE_URL/auth/register)

echo "   Response: $duplicate_response"
echo ""

# Test 5: Login with correct credentials
echo "5. Login with Correct Credentials"
echo "   POST $BASE_URL/auth/login"
login_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "admin123"
  }' \
  $BASE_URL/auth/login)

echo "   Response: $login_response"
login_token=$(echo $login_response | jq -r '.token // empty')
echo ""

# Test 6: Login with wrong password (should fail)
echo "6. Login with Wrong Password (should fail)"
echo "   POST $BASE_URL/auth/login"
wrong_pass_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "wrongpassword"
  }' \
  $BASE_URL/auth/login)

echo "   Response: $wrong_pass_response"
echo ""

# Test 7: Get current user info (protected endpoint)
echo "7. Get Current User Info (Protected)"
echo "   GET $BASE_URL/auth/me"
if [ -n "$admin_token" ]; then
  me_response=$(curl -s -H "Authorization: Bearer $admin_token" $BASE_URL/auth/me)
  echo "   Response: $me_response"
else
  echo "   Skipped (no token available)"
fi
echo ""

# Test 8: Access protected endpoint without token (should fail)
echo "8. Access Protected Endpoint Without Token (should fail)"
echo "   GET $BASE_URL/auth/me"
no_token_response=$(curl -s $BASE_URL/auth/me)
echo "   Response: $no_token_response"
echo ""

# Test 9: Invalid email format
echo "9. Register with Invalid Email (should fail)"
echo "   POST $BASE_URL/auth/register"
invalid_email_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "not-an-email",
    "password": "password123",
    "displayName": "Invalid Email",
    "role": "admin"
  }' \
  $BASE_URL/auth/register)

echo "   Response: $invalid_email_response"
echo ""

# Test 10: Weak password
echo "10. Register with Weak Password (should fail)"
echo "    POST $BASE_URL/auth/register"
weak_pass_response=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "weak@learng.com",
    "password": "short",
    "displayName": "Weak Password",
    "role": "admin"
  }' \
  $BASE_URL/auth/register)

echo "    Response: $weak_pass_response"
echo ""

echo "===================================="
echo "✅ Authentication Tests Complete"
