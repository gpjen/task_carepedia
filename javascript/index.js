import { createInterface } from 'readline';

const rl = createInterface({
  input: process.stdin,
  output: process.stdout
});

class HospitalQueue {
    constructor() {
        this.queue = []
        this.order = 'default'
        this.nextGender = 'F'
    }

    isDuplicateQueue(MRnumber) {
        return Boolean(this.queue.find(val => val?.MRnumber === MRnumber))
    }

    handleIn(MRnumber, gender) {
        console.log("Handle IN",MRnumber, gender)
    }

    handleOut() {
        console.log("handle Out");
    }

    setRoundRobin() {
        this.order = "roundRobin"
    }

    setDefault() {
        this.order = "default"
    }
}

console.log(' ===[ Hospital Queue Application ]===');

const queue = new HospitalQueue();
