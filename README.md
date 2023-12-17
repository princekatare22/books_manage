<h1>Book_Manager</h1>

<h2>The book_manager is Book Management API where the following operations are implemented,</h2>
<ul>
<li> Create()  : if new books are available they can be inserted into the database with the name of the book,the price and its genre.</li>
<li> Update()  : the attribute corresponding to the  books can be updated to new values.</li>
<li> GetByID() : this allows the user to fetch the information about the books from the database.</li>
<li> Delete()  : if the book is no more available it can be removed from the database.</li>
</ul>

<h2>Run the project using the following steps</h2>
<ol>
<li> Download the project in your local machine --> https://github.com/princekatare22/books_manage.git</li>
<li> Update the DB_USER and DB_PASSWORD in configs\.env</li>
<li> In the terminal, run command <b> go mod tidy </b></li>
<li> Run the project using command <b> go run .\main.go </b></li> 
</ol>

<h2>Sequence Daigram</h2>
<img src = "https://github.com/princekatare22/books_manage/assets/75197980/e606a6fd-84e2-46a4-8d08-d52d5a824f02">

<h2>Postman Collections</h2>
<h3>Create()</h3>
<img src = https://github.com/princekatare22/books_manage/assets/75197980/900be348-e958-48d0-a3da-04183189e005">
<h3>GetByID()</h3>
<img src = "https://github.com/princekatare22/books_manage/assets/75197980/a8cfbb74-4ef5-4e0d-8b47-28f09602ed47">
<h3>Update()</h3>
<img src = "https://github.com/princekatare22/books_manage/assets/75197980/98ec39c7-561d-4d97-8cd4-4f74d8c51e65">
<h3>Delete()</h3>
<img src = "https://github.com/princekatare22/books_manage/assets/75197980/bb5fe54e-0843-43bb-8071-be917316747e">




