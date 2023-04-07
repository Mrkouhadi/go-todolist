const searchForm = document.querySelector(".searchForm");

const data = new FormData(searchForm)
// data.append("csrf_token", "{{.CSRFToken}}")
function fetchDt(){
    fetch("/search",{
        method:"post",
        body:data,
    })
    .then(d => d.json())
    .then(d =>console.log(d))
    .catch(err => console.log("errr: ", err))
}
searchForm.addEventListener("submit",fetchDt)