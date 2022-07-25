using System.Collections.Generic;
namespace MicroBreweryRecipesAPI.Models
{  
    public class IngredientCategory
    {
        public int Id {get;set;}
        public string IngredientCategoryName {get;set;}

        public List<Ingredient> Ingredients{get;set;}
    }
}