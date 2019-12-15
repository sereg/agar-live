// import Go from './wasm_exec.js';


export async function  init() {
    const go = new Go();
    // let result = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject);
    // await go.run(result.instance);
    let {instance, module} = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject);
    await go.run(instance);
}

export class Universe {

    async cycle() {
        await window.cycle();
    }

    async restart() {
        await window.restart();
    }
    /**
     * @returns {string}
     */
    async export() {
        return window.export();
    }
    /**
     * @param {string} text
     */
    async import(text) {
        await window.import(text);
    }
    /**
     * @param {number} countAnimal
     * @param {number} countPlant
     */
    async generate(countAnimal, countPlant) {
        await window.generate(countAnimal, countPlant);
    }
    async backward() {
        await window.backward();
    }
    /**
     * @param {string} params
     */
    async setSize(params) {
        await window.setSize(params);
    }
    /**
     * @param {number} x
     * @param {number} y
     * @returns {string}
     */
    async changePosition(x, y) {
        return window.changePosition(x, y);
    }
    /**
     * @param {string} data
     * @param {number} x
     * @param {number} y
     * @returns {string}
     */
    async addFromJSON(data, x, y) {
        return window.addFromJSON(data, x, y);
    }
}