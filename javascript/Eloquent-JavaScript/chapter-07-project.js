(function() {
    "use strict";
  
    let active = null;
  
    const places = {
      "Alice's House": { x: 279, y: 100 },
      "Bob's House": { x: 295, y: 203 },
      Cabin: { x: 372, y: 67 },
      "Daria's House": { x: 183, y: 285 },
      "Ernie's House": { x: 50, y: 283 },
      Farm: { x: 36, y: 118 },
      "Grete's House": { x: 35, y: 187 },
      Marketplace: { x: 162, y: 110 },
      "Post Office": { x: 205, y: 57 },
      Shop: { x: 137, y: 212 },
      "Town Hall": { x: 202, y: 213 }
    };
    const placeKeys = Object.keys(places);
  
    const speed = 2;
  
    class Animation {
      constructor(worldState, robot, robotState) {
        this.worldState = worldState;
        this.robot = robot;
        this.robotState = robotState;
        this.turn = 0;
  
        let outer = window.__sandbox
            ? window.__sandbox.output.div
            : document.body,
          doc = outer.ownerDocument;
        this.node = outer.appendChild(doc.createElement("div"));
        this.node.style.cssText =
          "position: relative; line-height: 0.1; margin-left: 10px";
        this.map = this.node.appendChild(doc.createElement("img"));
        this.map.src = "http://eloquentjavascript.net/img/village2x.png";
        this.map.style.cssText = "vertical-align: -8px";
        this.robotElt = this.node.appendChild(doc.createElement("div"));
        this.robotElt.style.cssText = `position: absolute; transition: left ${0.8 /
          speed}s, top ${0.8 / speed}s;`;
        let robotPic = this.robotElt.appendChild(doc.createElement("img"));
        robotPic.src = "http://eloquentjavascript.net/img/robot_moving2x.gif";
        this.parcels = [];
  
        this.text = this.node.appendChild(doc.createElement("span"));
        this.button = this.node.appendChild(doc.createElement("button"));
        this.button.style.cssText =
          "color: white; background: #28b; border: none; border-radius: 2px; padding: 2px 5px; line-height: 1.1; font-family: sans-serif; font-size: 80%";
        this.button.textContent = "Stop";
  
        this.button.addEventListener("click", () => this.clicked());
        this.schedule();
  
        this.updateView();
        this.updateParcels();
  
        this.robotElt.addEventListener("transitionend", () =>
          this.updateParcels()
        );
      }
  
      updateView() {
        let pos = places[this.worldState.place];
        this.robotElt.style.top = pos.y - 38 + "px";
        this.robotElt.style.left = pos.x - 16 + "px";
  
        this.text.textContent = ` Turn ${this.turn} `;
      }
  
      updateParcels() {
        while (this.parcels.length) this.parcels.pop().remove();
        let heights = {};
        for (let { place, address } of this.worldState.parcels) {
          let height = heights[place] || (heights[place] = 0);
          heights[place] += 14;
          let node = document.createElement("div");
          let offset = placeKeys.indexOf(address) * 16;
          node.style.cssText =
            "position: absolute; height: 16px; width: 16px; background-image: url(http://eloquentjavascript.net/img/parcel2x.png); background-position: 0 -" +
            offset +
            "px";
          if (place == this.worldState.place) {
            node.style.left = "25px";
            node.style.bottom = 20 + height + "px";
            this.robotElt.appendChild(node);
          } else {
            let pos = places[place];
            node.style.left = pos.x - 5 + "px";
            node.style.top = pos.y - 10 - height + "px";
            this.node.appendChild(node);
          }
          this.parcels.push(node);
        }
      }
  
      tick() {
        let { direction, memory } = this.robot(this.worldState, this.robotState);
        this.worldState = this.worldState.move(direction);
        this.robotState = memory;
        this.turn++;
        this.updateView();
        if (this.worldState.parcels.length == 0) {
          this.button.remove();
          this.text.textContent = `Finished after ${this.turn} turns`;
          this.robotElt.firstChild.src =
            "http://eloquentjavascript.net/img/robot_idle2x.png";
        } else {
          this.schedule();
        }
      }
  
      schedule() {
        this.timeout = setTimeout(() => this.tick(), 1000 / speed);
      }
  
      clicked() {
        if (this.timeout == null) {
          this.schedule();
          this.button.textContent = "Stop";
          this.robotElt.firstChild.src =
            "http://eloquentjavascript.net/img/robot_moving2x.gif";
        } else {
          clearTimeout(this.timeout);
          this.timeout = null;
          this.button.textContent = "Start";
          this.robotElt.firstChild.src =
            "http://eloquentjavascript.net/img/robot_idle2x.png";
        }
      }
    }
  
    window.runRobotAnimation = function(worldState, robot, robotState) {
      if (active && active.timeout != null) clearTimeout(active.timeout);
      active = new Animation(worldState, robot, robotState);
    };
  })();
  
  //END OF VISUAL MAP AND ROBOT CODE
  
  
  // THESE ARE THE VARIABLES AND FUNCTIONS GIVEN IN THE CHAPTER (with my explanations)

var roads = [
    "Alice's House-Bob's House",
    "Alice's House-Cabin",
    "Alice's House-Post Office",
    "Bob's House-Town Hall",
    "Daria's House-Ernie's House",
    "Daria's House-Town Hall",
    "Ernie's House-Grete's House",
    "Grete's House-Farm",
    "Grete's House-Shop",
    "Marketplace-Farm",
    "Marketplace-Post Office",
    "Marketplace-Shop",
    "Marketplace-Town Hall",
    "Shop-Town Hall"
  ];
  
  function buildGraph(edges) {            //edges input is roads variable above
    let graph = Object.create(null);      //empty object with no inherited properties as opposed to {} which inherits Object.prototype
    function addEdge(from, to) {        //from and to taken from var roads are inputs (see for loop below)
      if (graph[from] == null) {        //if the graph doesn't have the property "from" it creates the property "to"
        graph[from] = [to]; 
      } else {                        //if the graph has the property "from" it pushes "to"
        graph[from].push(to);
      }
    }
    for (let [from, to] of edges.map(r => r.split("-"))) { //uses .map to create array and .split to separate "Start-End" format to [from, to] and iterates through roads variable
      addEdge(from, to);             //runs addEdge functions with parameters switched to iterate through every connection 
      addEdge(to, from);            // property "from" ends up with "to" as a value,  and property "to" ends up with "from" as a value
    }
    return graph;
  }
  
  var roadGraph = buildGraph(roads);  //console.log(roadGraph); //this is object with each .place as property with value = [] places it touches
  
  var VillageState = class VillageState {  
    constructor(place, parcels) {         //simply the constructor which gets called at the end of VillageState.random initally
      this.place = place;                 //and gets called at the end of state.move() 
      this.parcels = parcels;
    }
    move(destination) {
      if (!roadGraph[this.place].includes(destination)) {   
        return this;           // checks whether there is a road going from the current place to the destination, and if not, it returns the old state since this is not a valid move
      } else {
        let parcels = this.parcels
          .map(p => {   
            if (p.place != this.place) return p;    // I don't know why this line exists and it works without it
            return { place: destination, address: p.address };  //this maps a new array of objects(parcels) with the place = destination(where the robot is moving to)
          })
          .filter(p => p.place != p.address);               // filters out objects where the place and address are equal (dropping off parcels)
        return new VillageState(destination, parcels);       // returns new class (state) with updated destination (state.place) and parcels (state.parcels)
      }
    }
  };
  
  function runRobot(state, robot, memory) {
    for (let turn = 0; ; turn++) {            //iterates forever until there are no objects in parcels array
      if (state.parcels.length == 0) {
          console.log(`Done in ${turn} turns.`);
        //return turn;            //THIS WAS ADDED AS PART OF THE MEASURE A ROBOT SOLUTION
        break;
      }
      let action = robot(state, memory);
      state = state.move(action.direction);   // robot's .direction and .memory are generated by type of robot function, input into VillageState.move()
      memory = action.memory;             // robot's .memory is mailRoute or findRoute()
      console.log(`Moved to ${action.direction}`);
      //console.log(action);    //object with direction and memory, unless its randomRobot which has no memory
      //console.log(memory);
      //console.log(state);   // object with properties parcels (array of objects) and place (string)
    }
  }
  
  function randomPick(array) {
    let choice = Math.floor(Math.random() * array.length);      //this picks a random value within an array
    return array[choice];
  }
  
  function randomRobot(state) {
    return { direction: randomPick(roadGraph[state.place]) };   // chooses random value of current state.place property in roadGraph
  }
  
  //THIS IS THE STARTING STATE
  VillageState.random = function(parcelCount = 5) {         //the map starts with # parcels
    let parcels = [];
    for (let i = 0; i < parcelCount; i++) {
      let address = randomPick(Object.keys(roadGraph));   //sets random Object.keys of roadGraph (location) as parcel address
      let place;
      do {
        place = randomPick(Object.keys(roadGraph));       //sets random Object.keys of roadGraph (location) as parcel place
      } while (place == address);                         //loops so place and address don't start out the same
      parcels.push({ place, address });                   //adds object with place and address properties to array parcels
    }
    return new VillageState("Post Office", parcels);      //robot starts at Post Office, parcels start in random locations
  };
  
  var mailRoute = [
    "Alice's House",
    "Cabin",
    "Alice's House",
    "Bob's House",
    "Town Hall",
    "Daria's House",
    "Ernie's House",
    "Grete's House",
    "Shop",
    "Grete's House",
    "Farm",
    "Marketplace",
    "Post Office"
  ];
  
  function routeRobot(state, memory) {
    if (memory.length == 0) {
      memory = mailRoute;     //resets the memory to go around the route again
    }
    return { direction: memory[0], memory: memory.slice(1) };  // sets direction to next in memory array and removes that from the array
  }
  
  function findRoute(graph, from, to) {   //from is current robot.place, to is either parcels.place or parcels.address
    let work = [{ at: from, route: [] }];
    for (let i = 0; i < work.length; i++) {
      let { at, route } = work[i];
      for (let place of graph[at]) {      //iterates through values of roadGraph's property(robot's current location)
        if (place == to) return route.concat(place);    //if one of the values matches parcel's (either place or address) it's added to route
        if (!work.some(w => w.at == place)) {   // if no object in work array has .at value equal to place
          work.push({ at: place, route: route.concat(place) });  // add new object to work array 
        }
      } //console.log(work);
    } 
  }
  
  function goalOrientedRobot({ place, parcels }, route) {   //in runRobot function, robot (state, memory), {place, parcels} = state, route = memory
    if (route.length == 0) {
      let parcel = parcels[0];    //starts with first parcel robot has
      if (parcel.place != place) {      //if the parcel was not just picked up at the current place
        route = findRoute(roadGraph, place, parcel.place);    //run findRoute with to = parcel's current place
      } else {
        route = findRoute(roadGraph, place, parcel.address);  // if the parcel's current place is robots place, findRoute with parcel's address (where its going)
      }
    } //console.log(route);
    return { direction: route[0], memory: route.slice(1) }; //direction = first value in roadGraph's from(robot's current location), memory = removes first item from the array. when array is empty above code executes.
  }
  
//RUNS ANIMATION
//runRobotAnimation(VillageState.random(), randomRobot);
//runRobotAnimation(VillageState.random(), routeRobot, []);
runRobotAnimation(VillageState.random(), goalOrientedRobot, []);  

//RUNS IN CONSOLE
  //runRobot(VillageState.random(), randomRobot); 
  //runRobot(VillageState.random(), routeRobot, []);
  runRobot(VillageState.random(), goalOrientedRobot, []); 
  
  //console.log(roadGraph)
  
