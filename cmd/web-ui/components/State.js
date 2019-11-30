import React, {useState} from "react";
import {Box} from "@material-ui/core";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import useInterval from "./interval";

const stateURL = "http://localhost:8081/state";
const State = () => {


    const [st, setSt] = useState({
        latest_event_id: "1575083101660-0",
        users: {},
        items: {}
    });
    const fetchState = () =>
        fetch(stateURL)
            .then(res => res.json())
            .then(data => {
                setSt(data)
            });
    useInterval(() => {
        // Your custom logic here
        fetchState();
    }, 300);


    const users = st.users;
    const items = st.items;
    const keys = Object.keys(users);
    const itemKeys = Object.keys(items);

    return (
        <div>
            <Box>latest_event_id:{st.latest_event_id}</Box>
            User state:
            <Table aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>User ID </TableCell>
                        <TableCell align="left">Balance</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {keys.map(function (key) {
                        let row = users[key];
                        return (
                            <TableRow key={row.user_id}>
                                <TableCell align="left">{row.user_id}</TableCell>
                                <TableCell align="left">{row.balance}</TableCell>
                            </TableRow>
                        );
                    })}
                </TableBody>
            </Table>
            Items state:
            <Table aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>Item ID</TableCell>
                        <TableCell align="left">Price</TableCell>
                        <TableCell align="left">Remain</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {itemKeys.map(function (key) {
                        let row = items[key];
                        return (
                            <TableRow key={row.id}>
                                <TableCell align="left">{row.id}</TableCell>
                                <TableCell align="left">{row.price}</TableCell>
                                <TableCell align="left">{row.remain}</TableCell>
                            </TableRow>
                        );
                    })}
                </TableBody>
            </Table>
        </div>
    );
};

export default State;
