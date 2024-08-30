---
sidebar_position: 6
---

# DB

![item-code](/img/item-code.png)

1. The **item_code** is from SQL acc, it is use to map with the sku item on the platform so that when do posting it can correctly capture the item

![store](/img/store.png)

2. The red square is also from SQL acc, this is use to map the fee on platform so when do posting it will capture into the invoice

3. For the rest of the DB, it is not finalise, it just the first version of the general JSON, so please change it along the process

![statement](/img/statement.png)

4. The **fee_type** refer to the Order table's `revenue_amount` and `expenses_amount`

    1. One Order will have many statement id, this is to cater the lazada where payout is by statement

5. The **name** is the fee charge by the platform like shipping fee, affiliate fee...

6. The **Orer_item** table is a bridge between **SKU** and **Order** as they are many to many relationship 

7. The **User** in the db can ignore it as using the **Keycloak** to manage the user

8. Below is the full image of the DB 
![DB-image](/img/db-E-commerce.drawio.png)



