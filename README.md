# **Golang Developer Technical Test – Normalize Product IDs** 

# **Project Structure**

``` bash
📂 root
├── 📂 internal
│   ├── 📂 model
│   │   ├── order.go                    # Defines the data structures for orders
│   ├── 📂 service
│   │   ├── normalize_order.go          # Implements the core logic for normalizing orders
│   │   ├── normalize_order_test.go     # Unit tests for order normalization
├── go.mod                              # Go module file
├── go.sum                              # Dependencies lock file
├── README.md                           # Documentation for the project
```
---  

# **Installation & Usage**
### **1. Clone the repository**
``` bash
git clone git@github.com:Kurumichan666/technical-skills.git
cd technical-skills
```

### **2. Run Tests** 
``` bash
go test ./...
```
--- 

## **Objective**  
This project is a **Technical Skills Test** for the **Golang Developer** position. The goal is to **convert and match Product IDs** from online platforms into a standardized **Internal Code** through the following process:

## **Processing Steps**  

### **1. Input Processing via `NormalizeOrder` Service** 
- The main function responsible for cleaning and transforming the input data.  

### **2. Extract Product IDs using `splitProductID` function**  
- This function separates Product IDs from platform-specific strings.  
- Since the provided test cases **are not highly complex**, **regex is avoided** to maintain better performance.  

#### **Example of ID Extraction**  
 
Input: " --FG0A-CLEAR-OPPOA3*2/%20xFG0A-MATTE-OPPOA3 "
Output: ["FG0A-CLEAR-OPPOA3*2", "FG0A-MATTE-OPPOA3"]

Input: "x2-3&FG0A-CLEAR-IPHONE16PROMAX"
Output: ["FG0A-CLEAR-IPHONE16PROMAX"]
```

#### **Limitations of `splitProductID`** 🚧  
- ❌ If the format is `{filmTypeID} - {textureID} - {phoneModelID}&{text}`, the function **may not extract the correct ID**.  
- 📌 To handle such cases, the logic for removing `"&"` would need to be improved to apply only under the correct conditions.  

---  

### **3. Extract Ordered Quantity via `extractQuantity` function**  
- Find the number of items ordered 

#### **Example Input & Output**  
```plaintext
Input: "FG0A-MATTE-IPHONE16PROMAX*3"
Output: productID: "FG0A-MATTE-IPHONE16PROMAX", quantity: 3

Input: "FG0A-MATTE-IPHONE16PROMAX"
Output: productID: "FG0A-MATTE-IPHONE16PROMAX", quantity: 1
```

---  

### **4. Extract Product Details via `splitProductDetail` function**  
- Parses the **Product ID** into its respective components.  

#### **Example Input & Output**  
```plaintext
Input: "FG0A-CLEAR-IPHONE16PROMAX"
Output: filmTypeID: "FG0A", textureID: "CLEAR", phoneModelID: "IPHONE16PROMAX"

Input: "FG0A-PRIVACY-IPHONE16PROMAX-B"
Output: filmTypeID: "FG0A", textureID: "PRIVACY", phoneModelID: "IPHONE16PROMAX-B"
```

---  

### **5. Calculate Pricing & Generate `CleanedOrder`**  
- Uses the extracted details to calculate the ordered quantity and price per unit.  

#### **Pricing Limitations** 🚧  
- ❌ The current implementation **assumes all products have the same price**.  
- 📌 If different products have **varying unit prices**, an additional lookup from a database or pricing table would be required.  

---  

### **6. Add giveaway Items (Wiping Cloth & Cleaners)**  
---
### **7. Generate the Final Output**  

--- 

## **Summary**  
This code efficiently **normalizes Product IDs** and **processes orders** for general use cases.  
However, **some limitations** exist that may require enhancements 
