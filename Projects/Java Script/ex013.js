var agora = new Date()
var diasem = agora.getDay();
console.log(`${diasem}`)
switch(diasem)
{
    case 0:
        console.log("hoje e domingo");
        break;
        case 1:
        console.log("hoje e segunda");
        break;
    case 2:
        console.log("hoje e terca");
        break;
    case 3:
        console.log("hoje e quarta");
        break;
    case 4:
        console.log("hoje e quinta");
        break;  
    case 5:
        console.log("hoje e sexta");
        break;  
    case 6:
        console.log("hoje e sabado");
        break;  
    default:
        console.log("dia invalido");
}