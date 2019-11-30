import React from "react";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import ListEvent from "../components/ListEvent";
import State from "../components/State";
import PublisEvent from "../components/PublistEvent";

const Home = () => (
    <div>
        <Grid container spacing={1}>
            <Grid item xs={6}>
                <Paper>
                    <h3>List Events:</h3>
                    <ListEvent key="list_event"/>
                    <h3>Current State:</h3>
                    <State key="current-state"/>
                </Paper>
            </Grid>
            <Grid item xs={6}>
                <Paper><PublisEvent/></Paper>
            </Grid>
        </Grid>
    </div>
);

export default Home;
