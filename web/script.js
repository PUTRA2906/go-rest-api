document.addEventListener("DOMContentLoaded", () => {
  // Fungsi untuk mengambil data buku dari API
  function fetchBooks() {
    fetch("http://localhost:8080/books") // URL API kamu
      .then((response) => {
        // Cek apakah response berhasil
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        console.log("Books data received:", data);
        const bookList = document.getElementById("book-list");
        bookList.innerHTML = ""; // Reset isi list

        data.forEach((book) => {
          const listItem = document.createElement("li");
          listItem.innerHTML = `
              <span class="book-title">${book.title}</span>
              <span class="book-author">${book.author}</span>
            `;
          bookList.appendChild(listItem);
        });
      })
      .catch((error) => {
        console.error("Error fetching books:", error);
      });
  }

  // Panggil fungsi untuk mengambil buku saat halaman dimuat
  fetchBooks();
});
