package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var generatePermissions = &cobra.Command{
	Use: "groups",
	Short: "Generate list of users and group membership",
	Long: `Generate list of users and cognito userpool membership`,
	Run: listGroups,

}
func init(){
	rootCmd.AddCommand(generatePermissions)
	if err:= godotenv.Load(".env"); err != nil{
		log.Fatal("No .env file found", err)
	}
}

func listGroups(cmd *cobra.Command, args []string){
	// var
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
			// group magic
			groupsInput := cognito.AdminListGroupsForUserInput{
				UserPoolId: aws.String(usrPool),
				Username:   user.Username,
		}

		groupsResult, groupsErr := client.AdminListGroupsForUser(
			context.TODO(),
			&groupsInput)
		if groupsErr != nil {
				fmt.Println("Error retrieving group membership:", groupsErr)
				return
		}

		var groups []string
		for _, group := range groupsResult.Groups {
				groups = append(groups, *group.GroupName)
		}
			fmt.Printf("EMAIL: %s, groups: %v\n", email, groups)

		}
}