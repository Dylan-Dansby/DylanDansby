//Written by Dylan Dansby, July 26th, 2025
#include <stdio.h>

//Function that calculates velocity
float getVelocity(int p1, int p2, int time){
    if(time<=0){
        return 0;
    }
    if(p1<p2){
        return (p2-p1)/time;
    }
    if(p1>p2){
        return (p1-p2)/time;
    }else{
        return 0;
    } 
}

//Function that Calculates Direction
char* getDirection(int p1, int p2){
    if(p1>p2){
        return "west";
    } 
    if(p1<p2){
        return "east";
    } else {
        return "No Direction";
    }
    
}

//Function that calculates Kinetic Energy of Aircraft
double kEnergy(int mass, float velocity){
    return (0.5)*(double)mass*(velocity*velocity);
}

int main(){

    int mass,p1,p2,t = 0;
    printf("Please Input Aircraft Mass, point 1 and point 2 (in integer form), and the time in seconds it took to fly from 1 to 2.\n");
    //input
    scanf("%d, %d, %d, %d",&mass, &p1,&p2,&t);
    

    printf("%d,%d,%d,%d",mass,p1,p2,t);

    //variable assignments
    float velocity = getVelocity(p1,p2,t);
    char* direction = getDirection(p1,p2);
    double kineticEnergy = kEnergy(mass,getVelocity(p1,p2,t));
    float momentum = mass * velocity;

    //output
    printf("\nAircraft is flying %s at %.3g m/s. Kinetic Energy: %.5g Joules. Momentum: %.5g kg*m/s", direction,velocity,kineticEnergy,momentum);
    

    return 0;
}
