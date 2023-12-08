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
func Test(){
	var usrPool = os.Getenv("USRPOOL")
	sess, err := config.LoadDefaultConfig(context.TODO(),
							config.WithRegion("us-east-1"))
	client := cognito.NewFromConfig(sess)
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

		for i, user :=range result.Users{
			if &user == nil {
				continue
			}
			
			fmt.Printf("%d user %s created %v\n", i, *user.Username, user.UserCreateDate)

		}
}

func init(){
	rootCmd.AddCommand(generateList)
	if err:= godotenv.Load(".env"); err != nil{
		log.Fatal("No .env file found", err)
	}
}




func fetchList(cmd *cobra.Command, args []string){

	// add flags to customize output: default only emails
	var usrPool = os.Getenv("USRPOOL")
	cognitoClient := NewCognitoClient("us-east-1",os.Getenv("CLNTID"))

	result, err := cognitoClient.List(usrPool)
	if err != nil{
		panic(err)
	}
	for i, user :=range result.Users{
		if &user == nil {
			continue
		}
		fmt.Printf("%d user %s created %v\n", i, *user.Username, user.UserCreateDate)
	}
}

type CognitoClient interface {
	LogIn(email string, password string) (error, string)
	List(userPool string)(*cognito.ListUsersOutput, error)
}

type awsCognitoClient struct {
	cognitoClient *cognito.Client
	appClientId string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) CognitoClient{
	sess, err := config.LoadDefaultConfig(context.TODO(),
							config.WithRegion(cognitoRegion))
	client := cognito.NewFromConfig(sess)
	if err != nil {
		panic((err))
	}
	return &awsCognitoClient{
		cognitoClient: client,
		appClientId: cognitoAppClientID,
	}
}

func (ctx *awsCognitoClient) LogIn(email string, password string) (error, string){
	return nil, fmt.Sprintf("%s, %s", email, password)
}

func (ctx *awsCognitoClient) List(userPool string)(*cognito.ListUsersOutput, error){
	result, err := ctx.cognitoClient.ListUsers(
		context.TODO(),
		&cognito.ListUsersInput{
			UserPoolId: &userPool,
		}	)
		if err != nil {
			log.Printf("Couldn't list users. Here's why: %v\n", err)
			return nil, err
		}
	return result, err

}
