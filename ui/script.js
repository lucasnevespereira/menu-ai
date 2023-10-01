const generateBtn = document.getElementById("generate-btn");
const output = document.getElementById("output");
generateBtn.addEventListener("click", async () => {
    let prompt = "tocomplete"
    const content = await fetchGenerateMenu(prompt)
    output.innerHTML = content
})
async function fetchGenerateMenu(prompt) {
    try {
        const response = await fetch("/generate-menu", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ prompt }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.error || "An error occurred";
            throw new Error(errorMessage);
        }
    
        const data = await response.json();
        return data.completion;
      } catch (e) {
        console.error(e.message);
        return "An error occurred. Please try again later";
      }
}

