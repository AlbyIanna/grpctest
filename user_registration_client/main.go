/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	pb "grpctest/user_registration"

	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	testName        = "testUser"
	testEmail       = "test@test.com"
	testPassword    = "testPassword"
	testBirthday    = 723228323
	testPhoneNumber = "+39 3469103322"
	testQuestion    = "How did your first dog call you"
	testAnswer      = "peter"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserRegistrationClient(conn)

	name := testName
	email := testEmail
	password := testPassword
	birthday := &timestamp.Timestamp{Seconds: testBirthday}
	phone_number := testPhoneNumber
	question := testQuestion
	answer := testAnswer

	for i := 1; i < len(os.Args); i++ {
		switch i {
		case 1:
			name = os.Args[1]
		case 2:
			email = os.Args[2]
		case 3:
			password = os.Args[3]
		case 4:
			input_birthday, _ := strconv.ParseInt(os.Args[4], 0, 64)
			birthday = &timestamp.Timestamp{Seconds: input_birthday}
		case 5:
			phone_number = os.Args[5]
		case 6:
			question = os.Args[6]
		case 7:
			answer = os.Args[7]
		default:
		}
	}
	security_question := &pb.RegisterUserRequest_SecurityQuestion{Question: question, Answer: answer}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{Email: email, Username: name, Password: password, Birthday: birthday, PhoneNumber: phone_number, SecurityQuestion: security_question})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Hello user %s, your account has been set up correctly with e-Mail: %s", r.Username, r.Email)
}
