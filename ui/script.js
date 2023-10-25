const generateBtn = document.getElementById("generate-btn");
const loader = document.getElementById("loader");
const menuContentDiv = document.getElementById("menu-content");


generateBtn.addEventListener("click", generateMenu);

async function generateMenu() {
    loader.style.display = "block";
    const prompt = getAIPrompt()
    const menuContent = await getAIMenuContent(prompt)
    const formattedContent = extractAndFormatMenuContent(menuContent);
    menuContentDiv.innerHTML = formattedContent;
    loader.style.display = "none";
}

async function getAIMenuContent(prompt) {
    try {
        const response = await fetch("/generate-menu", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({prompt}),
        });

        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        const data = await response.json();
        return data.completion;
    } catch (error) {
        console.error(error);
        return "An error occurred. Please try again later.";
    }
}

function getAIPrompt() {
    const maxCalories = parseFloat(document.getElementById("max-calories").value);
    const maxCarbs = parseFloat(document.getElementById("max-carbs").value);
    const maxProteins = parseFloat(document.getElementById("max-proteins").value);
    const maxFats = parseFloat(document.getElementById("max-fats").value);
    const allergies = Array.from(
        document.getElementById("allergies").selectedOptions
    ).map((option) => option.value);
    return `Je veux créer un menu avec les précisions suivantes:
    
Calories max par jour: ${maxCalories}
Glucides max par jour: ${maxCarbs}%
Protéines max par jour: ${maxProteins}%
Lipides max par jour: ${maxFats}%

Allergies: ${allergies.join(", ")}

Préciser les quantités à utiliser (en grammes) pour chaque ingredient

Merci de separer le début du menu par START et la fin par END`;
}

function extractAndFormatMenuContent(menuContent) {
    const startIdx = menuContent.indexOf('START');
    const endIdx = menuContent.indexOf('END');

    if (startIdx !== -1 && endIdx !== -1) {
        const contentBetweenStartAndEnd = menuContent.slice(startIdx + 6, endIdx).trim();
        const formattedContent = contentBetweenStartAndEnd
            .replace(/-\s+/g, '• ') // Replace bullet points
            .replace(/:\n/g, ':\n\n') // Add extra newline after colons
            .replace(/\n\n/g, '<br>') // Convert double newline to HTML line break
            .replace(/\n/g, '<br>') // Convert remaining single newline to HTML line break;

        return formattedContent;
    }

    return '';
}

