import React, {useState} from "react";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import useInterval from "./interval";

const eventsURL = "http://localhost:8081/events";
const rows = [
    {
        event: {
            id: "1575083870094-0",
            type: "TopUp",
            user_id: 10,
            amount: 100
        },
        err: null
    }
];
const ListEvent = () => {
    const [events, setEvent] = useState([]);
    const fetchEvents = () =>
        fetch(eventsURL)
            .then(res => res.json())
            .then(data => {
                setEvent(data)
            });
    useInterval(() => {
        // Your custom logic here
        fetchEvents();
    }, 300);
    return (
        <Table aria-label="simple table">
            <TableHead>
                <TableRow>
                    <TableCell>ID</TableCell>
                    <TableCell align="left">Data</TableCell>
                    <TableCell align="left">Err</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {events.map(row => (
                    <TableRow key={row.event.id}>
                        <TableCell align="left">{row.event.id}</TableCell>
                        <TableCell align="left">{JSON.stringify(row.event)}</TableCell>
                        <TableCell align="left">{row.err}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    );
};

export default ListEvent;
