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

var (
	username string
	groupName string
	includeGroup bool
)

var modifyPermissions = &cobra.Command{
	Use: "modify",
	Short: "Modify permissions for a user",
	Long: `Modify permissions for a user registered in a cognito pool
	Example usage: 
	ldc modify -u d107d39c-df84-4c5f-98ba-97a778507df3 -g NWERN -i=true (short-hand flags)
	ldc modify --username d107d39c-df84-4c5f-98ba-97a778507df3 --group NWERN --include=true (long-form flags)`,
	Run: addRemovePermissions,

}
func init(){
	rootCmd.AddCommand(modifyPermissions)
	rootCmd.PersistentFlags().StringVarP(&username, "username","u","", "Username of cognito userpool user.")
	rootCmd.PersistentFlags().StringVarP(&groupName, "group", "g", "", "Cognito Group Name")
	rootCmd.PersistentFlags().BoolVarP(&includeGroup, "include", "i", false, "Include user in group (true) or remove from group (false)")

	if err:= godotenv.Load(".env"); err != nil{
		log.Fatal("No .env file found", err)
	}
}

func addRemovePermissions(cmd *cobra.Command, args []string){
	
	var usrPool = os.Getenv("USRPOOL")
	sess, err := config.LoadDefaultConfig(context.TODO(),
							config.WithRegion("us-east-1"))
	client := cognito.NewFromConfig(sess)

	operation:= "AddToGroup"
	if !includeGroup {
		operation = "RemoveFromGroup"
	}

	addParams:= &cognito.AdminAddUserToGroupInput{
		GroupName: &groupName,
		UserPoolId: &usrPool,
		Username: &username,
	}

	removeParams:= &cognito.AdminRemoveUserFromGroupInput{
		GroupName: &groupName,
		UserPoolId: &usrPool,
		Username: &username,
	}
	

	if operation != "AddToGroup"{
		_, err = client.AdminRemoveUserFromGroup(context.TODO(), removeParams)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("User", username, "successfully removed from group", groupName)
	} else {
		_, err = client.AdminAddUserToGroup(context.TODO(), addParams)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("User", username, "successfully added to group", groupName)
	}
	
	
}