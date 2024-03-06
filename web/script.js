    function searchOrder() {
        const order_uid = document.getElementById('order_uid').value;
        fetch(`http://localhost:8080/orders/${order_uid}`)
            .then(response => response.json())
            .then(data => {
                document.getElementById('result').innerText = JSON.stringify(data, null, 2);
            })
            .catch(error => {
                document.getElementById('result').innerText = 'Ошибка: ' + error;
            });
    }


