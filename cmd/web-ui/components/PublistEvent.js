import React, {useState} from "react";
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';


const PublisEvent = () => {
    const topup = useFormInput({
        "type": "TopUp",
        "user_id": 7,
        "amount": 100
    });
    const addItem = useFormInput({
        "type": "AddItem",
        "item_id": "cpu",
        "count": 100
    });
    const order = useFormInput({
        "type": "Order",
        "user_id": 7,
        "item_ids": ["cpu"],
        "item_quantities": [2]
    });
    return (
        <div>
            <form noValidate autoComplete="off">
                <div>
                    <TextField
                        id="standard-multiline-flexible"
                        label="Top-up"
                        multiline
                        rowsMax="10"
                        value={JSON.stringify(topup.value, null, "\t")}
                        onChange={topup.onChange}
                        margin="normal"
                    />
                </div>
                <Button variant="contained" color="primary" onClick={() => {
                    publishEvent(topup.value)
                }}>
                    Top-up
                </Button>
            </form>
            <form noValidate autoComplete="off">
                <div>
                    <TextField
                        id="standard-multiline-flexible"
                        label="Add Items"
                        multiline
                        rowsMax="10"
                        value={JSON.stringify(addItem.value, null, "\t")}
                        onChange={addItem.onChange}
                        margin="normal"
                    />
                </div>
                <Button variant="contained" color="primary" onClick={() => {
                    publishEvent(addItem.value)
                }}>
                    Add Items
                </Button>
            </form>
            <form noValidate autoComplete="off">
                <div>
                    <TextField
                        id="standard-multiline-flexible"
                        label="Make Order"
                        multiline
                        rowsMax="10"
                        value={JSON.stringify(order.value, null, "\t")}
                        onChange={order.onChange}
                        margin="normal"
                    />
                </div>
                <Button variant="contained" color="primary" onClick={() => {
                    publishEvent(order.value)
                }}>
                    Make Order
                </Button>
            </form>
        </div>
    )
};

async function publishEvent(event) {
    let url = "http://localhost:8080/event/publish";
    let response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(event)
    });

    if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        console.log(json)
    } else {
        alert("HTTP-Error: " + response.status);
    }
}

function useFormInput(initialValue) {
    const [value, setValue] = useState(initialValue);

    function handleChange(e) {
        setValue(JSON.parse(e.target.value));
    }

    return {
        value: value,
        onChange: handleChange,
    }
}

export default PublisEvent;