package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Representation struct {
	ID              string
	Name            string
	Address         string
	PhoneNumber     string
	EmployeeNumbers int
	JoinDate        time.Time
}
const(
	ID_REP int = iota 
	NAME         
	ADDRESS         
	PHONENUMBER     
	EMPLOYEENUMBERS 
	JOINDATE        
)

type Region struct {
	Name            string
	Representations []Representation
}

func main() {
		region := Region{Name: "tehran"}
		// region := flag.String("region", "Tehran", "region")
		command := flag.String("command", "", "Order")
		flag.Parse()

		switch *command {
		case "list":
			records, _ := loadData("data.csv")
			fmt.Println(records[1:])
		case "get":
			str ,_ , err := region.get()
			if err != nil{
				fmt.Printf("Error, %s\n", err)
				break
			}
			fmt.Println("result:\n", str)
		case "create":
			region.create()
		case "edit":
			err := region.edit()
			if err != nil{
				fmt.Printf("Error, %s\n", err)
			}
		case "status":
			region.status()
		case "exit":
			return	
		}



}

func (region Region) get() ([]string, int , error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter ID:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	records, err := loadData("data.csv")
	if err != nil{
		return []string{}, -1, err
	}
	for	i := 1; i < len(records); i++ {
		if records[i][ID_REP] == id{
			return records[i], i, nil
		}
	}

	return []string{}, -1, fmt.Errorf("Representation is not found.")
}

func (region Region) create() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter ID:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("enter name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("enter address:")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Print("enter phoneNumber:")
	phoneNumber, _ := reader.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	fmt.Print("enter employeeNumber:")
	employeeNumber, _ := reader.ReadString('\n')
	employeeNumber = strings.TrimSpace(employeeNumber)
	employeeNumberInt, _ := strconv.Atoi(employeeNumber)

	repr := Representation{
		ID:              id,
		Name:            name,
		Address:         address,
		PhoneNumber:     phoneNumber,
		EmployeeNumbers: employeeNumberInt,
		JoinDate:        time.Now(),
	}

	region.Representations = append(region.Representations, repr)
	appendData("data.csv", repr)
	fmt.Println("Representation is created.")
}

func (region Region) edit() error{
	record, i, err := region.get()
	if err != nil {
		return err
	}

	records, err := loadData("data.csv")
	if err != nil {
		return err
	}



	fmt.Println("Record found. Which field would you like to edit?")
	fmt.Println("[1] Name\n[2] Address\n\n[3] Phone Number\n[4] Employee Count\n[5] Cancel\nPlease enter the option number:")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice{
	case "1":
		fmt.Print("Enter Name:")
		edValue, _ := reader.ReadString('\n')
		edValue = strings.TrimSpace(edValue)
		record[NAME] = edValue
		records[i] = record
		return saveData("data.csv", records)
	case "2":
		fmt.Print("Enter Address:")
		address, _ := reader.ReadString('\n')
		address = strings.TrimSpace(address)
		record[ADDRESS] = address
		records[i] = record
		return saveData("data.csv", records)
		
	case "3":
		fmt.Print("Enter Phone Number:")
		phoneNumber, _ := reader.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)
		record[PHONENUMBER] = phoneNumber
		records[i] = record
		return saveData("data.csv", records)

	case "4":
		fmt.Print("Enter Employee Count:")
		employeeCount, _ := reader.ReadString('\n')
		employeeCount = strings.TrimSpace(employeeCount)
		record[EMPLOYEENUMBERS] = employeeCount	
		records[i] = record
		return saveData("data.csv", records)

	case "5": 
		break
	default:
		fmt.Println("Invalid option.")
	}
	return nil
}

func(region Region) status() {
	records, _ := loadData("data.csv")
	var count int
	
	c := len(records)
	if c>0 {
		c--
	}
	fmt.Printf("number of representation: %v\n", c)


	for i, _ := range records{
		if i==0{
			continue
		}
		fmt.Printf("representation name: %s,\t count: %v\n", records[i][NAME],records[i][EMPLOYEENUMBERS])
		temp, _ := strconv.Atoi(records[i][EMPLOYEENUMBERS])
		count+=temp
	}
	fmt.Printf("number of employees: %v\n", count)

	
}

func appendData(filename string, representation Representation) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		representation.ID,
		representation.Name,
		representation.Address,
		representation.PhoneNumber,
		strconv.Itoa(representation.EmployeeNumbers),
		representation.JoinDate.Format("2006-01-02"),
	}
	return writer.Write(record)
}
func loadData(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return [][]string{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	header := []string{"id", "name", "address", "phone_number", "employee_numbers", "join_date"}
	if len(records)>0{
		if !slices.Equal(header, records[0]){
			temp := make([][]string, 0)
			temp[0] = header
			temp = append(temp[:1], records...)
			return temp, nil
		}
	}
	
	return records, nil
}

func saveData(filename string, records [][]string) error{
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	return writer.WriteAll(records)
}
