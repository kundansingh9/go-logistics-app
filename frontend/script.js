async function checkShipment() {
    const result = document.getElementById("result");

    result.innerText = "Checking service...";

    try {
        const response = await fetch("/shipment");
        const data = await response.text();

        result.innerText = "✅ " + data;
    } catch (error) {
        result.innerText = "❌ Backend not running";
    }
}