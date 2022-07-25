namespace MicroBreweryRecipesAPI.Models
{
    public enum SI
    {
        MG,G,DAG,KG,T
    }
    public class Ingredient
    {
        public int Id{get;set;}
        public string IngredientName{get;set;}
        public SI IngredientUnit{get;set;}
        public float IngredientQuantity{get;set;}
        public string IngredientDescription{get;set;}
    }
}