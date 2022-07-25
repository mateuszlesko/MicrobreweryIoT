namespace MicroBreweryRecipesAPI.Models
{
    public class RecipeIngredientList
    {
        //entity properties
        public int Id {get;set;}
        public SI RecipeIngredientUnit{get;set;}
        public float IngredientQuantity{get;set;}
        
        //relationships
        public int RecipeId {get;set;}
        public Recipe Recipe {get;set;}
        public int IngredientId {get;set;}
        public Ingredient Ingredient {get;set;}
    }
}