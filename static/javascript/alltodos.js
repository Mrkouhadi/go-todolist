const container = document.querySelector(".todolist");

// fetching JSON data from /alltodos
const fetchDATA= async ()=>{
        await fetch("/alltodos-json")
          .then(d => d.json())
          .then(d => {
            if (d.ok){
              d.todos.forEach(todo => {
                const box = document.createElement("div")
                const p = CreateEl("p", todo.Title)
                const e = CreateEl("p", todo.Email)
                const c = CreateEl("p", todo.Content)
                const d = CreateEl("p", todo.IsDone)
                
                box.append(p)
                box.append(e)
                box.append(c)
                box.append(d)
                box.style.marginBottom = "40px"
                box.style.padding = "10px"

                box.style.backgroundColor = "white"

                container.append(box)
              });
            }
          })
          .catch(err => console.log("errr: ", err))
}
fetchDATA()

// create an html element
const CreateEl=(ell, txt)=>{
  const el = document.createElement(ell)
  el.append(txt)
  el.style.padding = "10px"
  el.style.margin = "4px"

  el.style.backgroundColor = "pink"
  return el
}