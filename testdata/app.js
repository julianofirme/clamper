console.log('Running Node.js application');

function consumeCPU() {
  const start = Date.now();
  while (Date.now() - start < 100) {
  }
  setImmediate(consumeCPU);
}

let memoryConsumption = [];
function consumeMemory() {
  const size = 50 * 1024 * 1024; 
  const buffer = Buffer.alloc(size, 'a');
  memoryConsumption.push(buffer);
  console.log(`Allocated ${memoryConsumption.length * 50} MB of memory`);
  setTimeout(consumeMemory, 1000); 
}

consumeCPU();
consumeMemory();
