import { createInterface } from "readline";

const rl = createInterface({
  input: process.stdin,
  output: process.stdout,
});

class HospitalQueue {
  constructor() {
    this.queue = [];
    this.order = "default";
    this.nextGender = "F";
  }

  isDuplicateQueue(mrNumber) {
    return this.queue.some((patient) => patient.mrNumber === mrNumber);
  }

  handleIn(mrNumber, gender) {
    if (this.isDuplicateQueue(mrNumber)) {
      console.log(`error: patient with ${mrNumber} already in queue`);
      return;
    }
    this.queue.push({ mrNumber, gender });
  }

  handleOut() {
    if (this.queue.length < 1) {
      console.log("queue is empty.");
      return;
    }

    let deletePatient = {};
    if (this.order === "default") {
      deletePatient = this.queue.shift();
    } else if (this.order === "roundRobin") {
      const index = this.queue.findIndex(
        (patient) => patient?.gender === this.nextGender
      );

      if (index !== -1) {
        deletePatient = this.queue[index];
        this.queue.splice(index, 1);
      } else {
        deletePatient = this.queue[0];
        this.queue.splice(0, 1);
      }

      this.nextGender = this.nextGender === "M" ? "F" : "M";
    }
    console.log(`send : ${deletePatient?.mrNumber} ${deletePatient?.gender}`);
  }

  setRoundRobin() {
    this.order = "roundRobin";
  }

  setDefault() {
    this.order = "default";
  }
}

console.log(" ===[ Hospital Queue Application ]===");
let queue = new HospitalQueue();

rl.setPrompt("> ");
rl.prompt();
rl.on("line", (comandLine) => {
  const input = comandLine.trim();
  const command = input.split(" ");

  if (command.length < 1) {
    rl.prompt();
    return;
  }

  switch (command[0].toUpperCase()) {
    case "IN":
      if (command.length < 3) {
        console.log("invalid In command.");
        break;
      }
      const mrNumber = command[1];
      const gender = command[2].toUpperCase();
      if (gender !== "M" && gender !== "F") {
        console.log("invalid gender.");
        break;
      }
      queue.handleIn(`MR${mrNumber}`, gender);
      break;

    case "OUT":
      queue.handleOut();
      break;
    case "ROUNDROBIN":
      queue.setRoundRobin();
      break;
    case "DEFAULT":
      queue.setDefault();
      break;
    case "EXIT":
      process.exit(0);
    default:
      console.log("invalid command line.");
      break;
  }

  rl.prompt();
});
