namespace MicroBreweryRecipesAPI.Models
{
    public class RecipeStage
    {
        public int Id {get;set;}
        public int OptimalTemperature{get;set;} // in Celcius
        public long MashTime {get;set;}

        public Recipe Recipe{get;set;}
    }
}