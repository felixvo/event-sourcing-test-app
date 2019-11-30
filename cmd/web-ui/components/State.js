import React from "react";
import { Box } from "@material-ui/core";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";

const st = {
  latest_event_id: "1575083101660-0",
  users: {
    "1": {
      user_id: 1,
      balance: 400
    },
    "10": {
      user_id: 10,
      balance: 100
    },
    "1318": {
      user_id: 1318,
      balance: 500
    },
    "1847": {
      user_id: 1847,
      balance: 500
    },
    "2081": {
      user_id: 2081,
      balance: 500
    },
    "2540": {
      user_id: 2540,
      balance: 500
    },
    "3300": {
      user_id: 3300,
      balance: 500
    },
    "4059": {
      user_id: 4059,
      balance: 500
    },
    "4425": {
      user_id: 4425,
      balance: 500
    },
    "456": {
      user_id: 456,
      balance: 500
    },
    "7887": {
      user_id: 7887,
      balance: 500
    },
    "8081": {
      user_id: 8081,
      balance: 500
    }
  },
  items: {
    cpu: {
      id: "cpu",
      price: 47,
      remain: 91
    },
    hdd: {
      id: "hdd",
      price: 81,
      remain: 152
    },
    ram: {
      id: "ram",
      price: 25,
      remain: 146
    },
    ssd: {
      id: "ssd",
      price: 81,
      remain: 264
    }
  }
};
const State = () => {
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
          {keys.map(function(key) {
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
          {itemKeys.map(function(key) {
            let row = items[key];
            return (
              <TableRow key={row.item_id}>
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
