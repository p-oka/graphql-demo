import React from "react";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

import enhance from "./enhance";

function TaskList(props) {
  const { tasks } = props;

  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>ID</TableCell>
          <TableCell>タスク名</TableCell>
        </TableRow>
      </TableHead>
      {tasks.map(task => (
        <TableRow>
          <TableCell>{task.id}</TableCell>
          <TableCell>{task.name}</TableCell>
        </TableRow>
      ))}
    </Table>
  );
}

export default enhance(TaskList);