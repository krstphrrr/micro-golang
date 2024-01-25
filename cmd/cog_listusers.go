package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var generateList = &cobra.Command{
	Use: "list",
	Short: "Generate a list of users registered",
	Long: `Generate a list of all users registered in the ldc api.`,
	Run: fetchList,

}

func init(){
	rootCmd.AddCommand(generateList)
	if err:= godotenv.Load(".env"); err != nil{
		log.Fatal("No .env file found", err)
	}
}




func fetchList(cmd *cobra.Command, args []string){
	// fetch list of users, emails, datecreated
	// leave the cmd and args array to enable flags

	// accessing user pool by configuring a session :
	var usrPool = os.Getenv("USRPOOL")
	sess, err := config.LoadDefaultConfig(context.TODO(),
							config.WithRegion("us-east-1"))
	client := cognito.NewFromConfig(sess)

	// listusers docs: https://aws.github.io/aws-sdk-go-v2/docs/
	input := &cognito.ListUsersInput{ 
						AttributesToGet: []string{"email"},
					UserPoolId: &usrPool,
				}
	result,err := client.ListUsers(
		context.TODO(),
		input,
			)
		if err != nil {
			log.Printf("Couldn't list users. Here's why: %v\n", err)
			return
		}

		for _ , user :=range result.Users{
			var email string
			for _, attr := range user.Attributes {
				
				if *attr.Name == "email" {
					email = *attr.Value
					break
				} 
			}
			
			fmt.Printf("USER: %s, EMAIL: %s, DATE_CREATED: %v\n", *user.Username, email, user.UserCreateDate)

		}
}
